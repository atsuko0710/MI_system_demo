package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn
}

// NewUser 创建一个新用户对象
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String() // 获取远程ip地址

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		Conn: conn,
	}

	// 启动监听当前 User channel 的 goroutine
	go user.ListenMassage()

	return user
}

// ListenMassage 监听当前用户 channel,一旦有消息就发送到客户端
func (this *User) ListenMassage() {
	for {
		msg := <-this.C
		this.Conn.Write([]byte(msg + "\n"))
	}
}
