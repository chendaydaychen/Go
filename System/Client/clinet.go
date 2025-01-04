package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	conn       net.Conn
	Name       string
	flag       int
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	// 连接服务器
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn

	return client
}

func (c *Client) menu() bool {
	fmt.Println("1. 公聊")
	fmt.Println("2. 私聊")
	fmt.Println("3. 更新用户名")
	fmt.Println("4. 退出")

	var key int
	fmt.Scanln(&key)

	if key > 0 && key < 5 {
		c.flag = key
		return true
	} else {
		fmt.Println(">>>> 输入非法，请重新输入...")
		return false
	}

}

func (c *Client) UpdataName() bool {
	fmt.Println(">>>> 更新用户名")
	fmt.Print(">>>> 请输入用户名：")
	fmt.Scanln(&c.Name)
	// 发送更新用户名请求
	sendMsg := "rename|" + c.Name + "\n"
	_, err := c.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(">>>> 更新用户名失败...")
		return false
	}
	return true
}
func (c *Client) PublicChat() {
	fmt.Println(">>>> 公聊")
	fmt.Print(">>>> 请输入聊天内容，exit退出：")
	var chatMsg string
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发送chatMsg到服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := c.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println(">>>> 发送消息失败...")
				break
			}
		}
		chatMsg = ""
		fmt.Print(">>>> 请输入聊天内容，exit退出：")
		fmt.Scanln(&chatMsg)

	}
}

func (c *Client) SelectUser() {
	sendMsg := "who\n"
	_, err := c.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(">>>> 发送消息失败...")
		return
	}
	fmt.Print(">>>> 请输入聊天对象，exit退出：")
}

func (c *Client) PrivateChat() {
	c.SelectUser()
	var remoteName string
	var chatMsg string
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Print(">>>> 请输入聊天内容，exit退出：")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				_, err := c.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println(">>>> 发送消息失败...")
				}
			}
			chatMsg = ""
			fmt.Print(">>>> 请输入聊天内容，exit退出：")
			fmt.Scanln(&chatMsg)
		}
		c.SelectUser()
		fmt.Scanln(&remoteName)
	}

}

func (c *Client) run() {
	for c.flag != 4 {
		for !c.menu() {
		}
		// 根据不同的选择，执行不同的业务
		switch c.flag {
		case 1:
			c.PublicChat()
		case 2:
			c.PrivateChat()
		case 3:
			c.UpdataName()
		}
	}
}

func (c *Client) DealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, c.conn)
}

func main() {

	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>> 连接服务器失败...")
		return
	}

	// 启动一个单独的goroutine来处理server的返回消息
	go client.DealResponse()

	fmt.Println(">>>> 连接服务器成功...")

	//启动客户端的监听
	client.run()
}
