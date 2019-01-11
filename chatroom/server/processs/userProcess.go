package processs

import (
	"net"
	"syst/chatroom/common/message"
	"encoding/json"
	"fmt"
	"syst/chatroom/server/utils"
)

type UserProcess struct {
	//字段
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//var mes message.Message
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ",err)
		return
	}
	//先声明一个　resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//再声明一个LoginResMes,并完成赋值
	var loginResMes message.LoginResMes


	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	}else{
		//不合法
		loginResMes.Code = 500 //500状态码，表示该用户不存在
		loginResMes.Error = "该用户不存在，请注册再使用..."
	}

	//序列化
	data,err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail",err)
		return
	}

	//４.将data赋值给 resMes.Data
	resMes.Data = string(data)

	//5.对resMes 进行序列化,准备发送
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail",err)
		return
	}

	//6.发送data,我们将其封装到writePkg函数
	//因为使用分层模式（mvc）,我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return nil
}