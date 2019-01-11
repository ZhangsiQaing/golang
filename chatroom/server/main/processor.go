package main

import (
	"net"
	"fmt"
	"syst/chatroom/common/message"
	"syst/chatroom/server/processs"
	"syst/chatroom/server/utils"
	"io"
)

//先创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes 函数
//功能:根据客户端发送消息种类不同，决定调用那个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error){
	//fmt.Println(json.Unmarshal(mes))
	switch mes.Type {
	case message.LoginMesType:
		//处理服务器登录
		//创建一个UserProcess实例
		up := &processs.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessLogin(mes)

		//处理登录逻辑
	case message.RegisterMesType:
		return
		//处理注册
	default:
		fmt.Println("消息不存在，无法处理...")
	}
	return nil
}


func (this *Processor) process2() (err error){
	//读客户端发送的信息
	//defer this.Conn.Close()

	//循环接收客户端发送的信息
	for {
		tf := &utils.Transfer{
			Conn:this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出，服务端也退出..")
				return err
			}else {
				fmt.Println("readPkg err = ", err)
				return err
			}
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			//fmt.Println("readPkg(conn) err=",err)
			return err
		}
		//fmt.Println("mes =",mes)
		//return err
	}
}