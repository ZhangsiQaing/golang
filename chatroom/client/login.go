package main

import (
	"net"
	"fmt"
	"syst/chatroom/common/message"
	"encoding/json"
)

func login(userId int,userPwd string) (err error) {

	//下一步就要开始定协议
	//fmt.Println(" userId = %d userPwd = %s",userId,userPwd)

	//return nil
	net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	//２.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//４.将loginMes　序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marsha1 loginmes err=", err)
		return
	}

	//５.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marsha1 mes err=", err)
		return
	}

	//7.到这个时候　这个就是我们要发送的数据
	//7.1宪法data的长度发送给服务器
	//conn.Write(len(data))
	//先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint64

}