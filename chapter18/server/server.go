package main

import (
	"net"
	"fmt"
	"io"
)

func process(conn net.Conn){
	//循环接收客户端发送的数据
	defer conn.Close()

	fmt.Printf("服务器在等待客户端发送信息\n" )
	for {
		//创建一个新的切片
		buf := make([]byte,1024)
		//conn.Read(buf)
		//1、等待客户端通过conn发送信息
		//2、如果客户端没有write[发送],那么协程就阻塞在这里

		n,err := conn.Read(buf)
		if err == io.EOF {
			fmt.Printf("服务器的Read err = %v",err)
			return
		}

		//3.显示客户端发送的内容到服务器的终端
		fmt.Printf("%v",string(buf[:n]))
		if(string(buf[:n]) == "exit"){
			return
		}
	}


}

func main(){

	listen,err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil{

		fmt.Println("listen error = ",err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端来连接...")
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("conn err = ",err)
		} else {
		   fmt.Printf("Accept() suc con=%v 客户端ip=%v\n",conn,conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
