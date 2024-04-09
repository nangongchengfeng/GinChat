package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("config init success")
}

func InitMySQL() error {
	//自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	fmt.Println(viper.GetString("mysql.dns"))
	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if err != nil {
		return fmt.Errorf("mysql init fail: %s \n", err)
	}
	fmt.Println("mysql init success")
	return nil
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		//Addr:         viper.GetString("redis.addr"),
		//Password:     viper.GetString("redis.password"),
		//DB:           viper.GetInt("redis.DB"),
		//PoolSize:     viper.GetInt("redis.poolSize"),
		//MinIdleConns: viper.GetInt("redis.minIdleConn"),
		Addr:     "192.168.102.20:6379",
		Password: "123456",
		DB:       0,
	})
	pong, err := Red.Ping().Result()
	if err != nil {
		fmt.Println("redis init fail: %s \n", err)
	} else {
		fmt.Println("redis init success:", pong)
	}

}

// 固定变量
const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(channel string, msg string) error {
	fmt.Println("Publishing...", msg)
	_, err := Red.Publish(channel, msg).Result()
	if err != nil {
		fmt.Println("Publish error:", err)
	}
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(channel string) string {
	if Red == nil {
		log.Println("Redis client is nil")
	} else {
		log.Println("Subscribing to channel:", channel)
	}
	fmt.Println("Subscribing to...", channel)
	sub := Red.Subscribe(channel)

	msg, _ := sub.ReceiveMessage()
	defer sub.Close()
	return msg.Payload
}
