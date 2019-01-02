package main

import (
	"net"
	"fmt"
)

func main(){

	listen,err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil{
		fmt.Println("listen error = ",err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端来连接...")
		conn,err := listen.Accept();
		if err != nil{
			fmt.Println("conn err = ",err);
		} else {
		    fmt.Printf("Accept() suc con=%v\n",conn)
		}

	}
}
