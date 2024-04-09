package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	gorm.Model
	FormId   uint   //发送者
	TargetId uint   //接受者
	Type     string //消息类型 群聊 私聊 广播
	Media    int    //消息类型 文字 图片 音频
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

// Node 是一个包含 WebSocket 连接、数据队列和群组集合的结构体。
type Node struct {
	Conn      *websocket.Conn // WebSocket 连接
	DataQueue chan []byte     // 数据队列，用于存储待发送的数据
	GroupSets set.Interface   // 群组集合，用于存储该节点所属的群组信息
}

// clientMap 是一个全局变量，用于存储所有连接的节点信息。
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// rwLock 是一个全局变量，用于对 clientMap 的读写操作进行加锁。
var rwLock sync.RWMutex

// Chat 处理聊天请求
func Chat(writer http.ResponseWriter, request *http.Request) {
	// 1、获取请求参数
	//token := query.Get("token")
	query := request.URL.Query()
	Id := query.Get("userid")
	// 2、根据请求参数查询数据库，获取用户信息
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//msgType := query.Get("type")
	//targetId := query.Get("targetId")
	// context := query.Get("context")
	isvalida := true // checktoken
	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	// 3、创建一个 Node 结构体，用于存储 WebSocket 连接和数据队列
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 用户关系

	//. userid 跟 node绑定 并加锁
	rwLock.Lock()
	clientMap[userId] = node
	rwLock.Unlock()
	// 4、启动一个 goroutine 用于处理该连接的读写和数据发送
	//5.完成发送逻辑
	go sendProc(node)
	//6.完成接受逻辑
	go recvProc(node)
	sendMsg(userId, []byte("欢迎进入聊天系统"))
}

// sendProc 函数用于从节点的 DataQueue 通道读取数据，并通过 WebSocket 连接发送数据。
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// recvProc 函数用于从节点的 WebSocket 连接中接收数据，并将数据广播给其他节点。
func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] <<<<< ", data)
	}
}

// udpsendChan 是一个缓冲通道，用于存储待广播的数据。
var udpsendChan chan []byte = make(chan []byte, 1024)

// broadMsg 函数将数据放入 udpsendChan 通道，以便广播给其他节点。
func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// 完成udp数据发送协程
func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 96, 19),
		Port: 30000,
	})
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-udpsendChan:
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
}

// 完成udp数据接收协程
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])

	}
}

func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case strconv.Itoa(1): //私信
		sendMsg(int64(msg.TargetId), data)
		// case 2: //群发
		// sendGroupMsg()
		// case 3://广播
		// sendAllMsg()
		//case 4:
		//
	}
}

func sendMsg(userId int64, msg []byte) {
	rwLock.RLock()
	node, ok := clientMap[userId]
	rwLock.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
