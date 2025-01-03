package main

import "net"

type User struct {
	Name    string
	Address string
	C       chan string
	conn    net.Conn

	server *Server
}

// NewUser 创建并初始化一个新的用户对象。
// 该函数接收一个网络连接对象作为参数，用于后续与用户进行通信。
// 返回值是一个指向User结构的指针，表示新创建的用户实例。
// 此函数的核心作用是根据连接信息初始化用户对象，使其能够通过给定的连接进行通信。
func NewUser(conn net.Conn, server *Server) *User {
	// 获取连接的远程地址，并将其用作用户的名称和地址。
	// 这样做是为了在用户层面上标识连接来源，便于后续的通信和管理。
	userAddr := conn.RemoteAddr().String()

	// 初始化User结构体实例。
	// Name和Address字段被设置为连接的远程地址字符串，以便于跟踪用户来源。
	// C字段是一个字符串通道，用于接收和发送消息。
	// conn字段是私有的，存储了用户连接，以便于后续的读写操作。
	user := &User{
		Name:    userAddr,
		Address: userAddr,
		C:       make(chan string),
		conn:    conn,
		server:  server,
	}

	// 创建一个 goroutine来监听用户的消息并将其写入通道 C。
	go user.ListenMessage()

	// 返回初始化后的用户实例。
	return user
}

// ListenMessage 监听用户消息。
//
// 该方法在一个无限循环中接收通过通道 C 发送的消息，并将其写入用户的连接 conn。
// 这使得用户能够持续接收和显示来自其他用户的消息或系统消息。
func (u *User) ListenMessage() {
	for {
		// 从通道 C 接收消息。
		msg := <-u.C

		// 将接收到的消息写入用户连接，确保用户能够看到消息。
		// 消息末尾添加换行符以改善可读性。
		u.conn.Write([]byte(msg + "\n"))
	}
}

func (u *User) Online() {
	//用户上线，加入在线列表
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()

	//广播上线消息
	u.server.Broadcast(u, "已上线")
}

func (u *User) Offline() {
	// 用户下线，删除在线列表
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()

	// 广播下线消息
	u.server.Broadcast(u, "已下线")
}

func (u *User) SendMsg(msg string) {
	u.conn.Write([]byte(msg))
}

func (u *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户
		u.server.mapLock.Lock()
		for _, user := range u.server.OnlineMap {
			onlineMsg := "[" + user.Address + "]" + user.Name + ":在线...\n"
			u.SendMsg(onlineMsg)
		}
		u.server.mapLock.Unlock()
		return
	} else {
		u.server.Broadcast(u, msg)
	}

}
