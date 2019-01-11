package main

import (
	"net"
	"fmt"
	//_ "syst/chatroom/common/message"
	//"syst/chatroom/server/utils"
	//_ "encoding/json"
	//"io"
	//"syst/chatroom/common/utils"
	//"io"
)


//func serverProcessLogin(conn net.Conn,mes *message.Message) (err error) {
//	//var mes message.Message
//	var loginMes message.LoginMes
//	json.Unmarshal([]byte(mes.Data),&loginMes)
//	if err != nil {
//		fmt.Println("json.Unmarshal fail err = ",err)
//		return
//	}
//	//先声明一个　resMes
//	var resMes message.Message
//	resMes.Type = message.LoginResMesType
//
//	//再声明一个LoginResMes,并完成赋值
//	var loginResMes message.LoginResMes
//
//
//	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
//		//合法
//		loginResMes.Code = 200
//	}else{
//		//不合法
//		loginResMes.Code = 500 //500状态码，表示该用户不存在
//		loginResMes.Error = "该用户不存在，请注册再使用..."
//	}
//
//	//序列化
//	data,err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("json.Marshal fail",err)
//		return
//	}
//
//	//４.将data赋值给 resMes.Data
//	resMes.Data = string(data)
//
//	//5.对resMes 进行序列化,准备发送
//	data,err = json.Marshal(resMes)
//	if err != nil {
//		fmt.Println("json.Marshal fail",err)
//		return
//	}
//
//	//6.发送data,我们将其封装到writePkg函数
//	err = utils.WritePkg(conn,data)
//	return nil
//}

////编写一个ServerProcessMes 函数
////功能:根据客户端发送消息种类不同，决定调用那个函数来处理
//func serverProcessMes(conn net.Conn,mes *message.Message) (err error){
//	//fmt.Println(json.Unmarshal(mes))
//	switch mes.Type {
//	   case message.LoginMesType:
//	   	//处理服务器登录
//	   	  err = serverProcessLogin(conn,mes)
//
//		   //处理登录逻辑
//		   case message.RegisterMesType:
//		   	return
//		//处理注册
//		default:
//		    fmt.Println("消息不存在，无法处理...")
//	}
//	return nil
// }




func process(conn net.Conn){
	//读客户端发送的信息
   defer conn.Close()

   //这里调用总控,创建一个
   processor := &Processor{
   	   Conn:conn,
   }
   err := processor.process2()
   if err != nil {
   	   fmt.Println("客户端和服务端通讯协程错误＝err",err)
   }
}

func main(){
	//提示
	fmt.Println("服务器[新结构]在8889端口监听....")
	listen,err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil{
		fmt.Println("listen err = ",err)
		return
	}


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