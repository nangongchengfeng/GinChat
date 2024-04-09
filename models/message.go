package models

import (
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
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
