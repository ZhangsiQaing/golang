package main

import (
	"net"
	"fmt"
)

func process(conn net.Conn){
	//读客户端发送的信息

}

func main(){
	listen,err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil{
		fmt.Println("listen err = ",err)
		return
	}
	defer listen.Close()

	for{
		fmt.Println("等待客户端来连接服务器")
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("accept err = ",err)
			return
		}

		//一旦连接成功,则启动一个协程和客户端保持通讯
		go process(conn)
	}

}