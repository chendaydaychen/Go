package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播
	Message chan string
}

// NewServer 创建并初始化一个新的Server实例。
// 参数:
//
//	ip: 服务器的IP地址。
//	port: 服务器的端口号。
//
// 返回值:
//
//	*Server: 指向新创建的Server实例的指针。
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
}

// ListenMessager 监听消息分发器
// 该方法用于监听服务器接收到的消息，并将这些消息分发给所有在线客户端
// 它利用一个无限循环来不断地检查是否有新消息到达，一旦有消息到达，就将其分发给所有在线客户端
func (s *Server) ListenMessager() {
	for {
		// 从服务器的消息通道中接收消息
		msg := <-s.Message

		// 在访问在线客户端映射表之前锁定映射表，以防止并发访问冲突
		s.mapLock.Lock()

		// 遍历在线客户端映射表，将接收到的消息发送给每个在线客户端
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}

		// 解锁映射表，允许其他协程访问
		s.mapLock.Unlock()
	}
}

// Broadcast 是一个服务器方法，用于向所有用户广播消息。
// 此方法接收一个用户实例和一条消息字符串作为参数，
// 组装这些信息，并将它们发送到消息通道，以便所有连接的用户都可以接收到。
//
// 参数:
//
//	user *User - 发送消息的用户信息，包括用户的地址和名称。
//	msg string - 用户发送的消息内容。
func (s *Server) Broadcast(user *User, msg string) {
	// 组装消息格式为："[用户地址]用户名:消息内容"
	sendMsg := "[" + user.Address + "]" + user.Name + ":" + msg
	// 将组装好的消息发送到服务器的消息通道
	s.Message <- sendMsg
}

// Handler 处理客户端连接
// 该函数主要负责处理客户端与服务器建立的连接
// 参数:
//
//	conn net.Conn: 代表客户端与服务器之间的连接
func (s *Server) Handler(conn net.Conn) {
	// 当连接建立成功时，打印提示信息
	//fmt.Println("连接建立成功")

	user := NewUser(conn, s)

	user.Online()

	//监听是否活跃
	isLive := make(chan bool)

	//读取并处理来自客户端的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			msg := string(buf[:n-1])

			user.DoMessage(msg)

			isLive <- true
		}
	}()
	//阻塞
	for {
		select {
		case <-isLive:
			//重置定时器
			//如果客户端发送了消息，则重置定时器，使其继续计时
			continue
		case <-time.After(time.Second * 10):
			//超时，强制关闭

			user.SendMsg("你被踢了")
			close(user.C)
			conn.Close()
			return
		}
	}
}

// Start 启动服务器并开始监听指定的IP和端口
// 该方法负责初始化服务器的监听器，接受客户端连接，并为每个连接分配处理程序
func (s *Server) Start() {
	// socket监听
	// 这里使用net.Listen创建一个TCP监听器，它将监听服务器的Ip和Port属性指定的地址
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	// 关闭连接
	// 在处理完所有请求后，确保每个连接都被正确关闭，以释放资源
	defer listener.Close()

	//启动监听Msessage方法
	go s.ListenMessager()

	for {
		// 接受连接
		// 接受监听器接收到的每个连接，并为每个连接创建一个新的goroutine来处理数据

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接失败:", err)
			continue
		}
		// 执行处理程序
		// 对于每个接受的连接，这里将调用一个处理函数来读取和响应客户端请求
		go s.Handler(conn)
	}

}
