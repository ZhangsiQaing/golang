package utils

import (
	"net"
	"syst/chatroom/common/message"
	"fmt"
	"encoding/binary"
	"encoding/json"
)

func ReadPkg(conn net.Conn) (mes message.Message,err error){
	buf := make([]byte,8096)
	fmt.Println("读取客户端发送的数据....")
	_,err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}

	//根据buf[:4] 转成一个　uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据 pkgLen 读取消息内容
	n,err := conn.Read(buf[:pkgLen])
	//fmt.Println(n,string(buf),err)
	if n != int(pkgLen) || err != nil {
		//fmt.Println("connn.Read fail err = ",err)
		//err = errors.New("read pkg header error")
		return
	}

	//buf反序列化成 -> message.Message
	err = json.Unmarshal(buf[:n],&mes)
	fmt.Println(err)
	if err != nil {
		//fmt.Println("json.Unmarshal fail err = ",err)
		//err = errors.New("read pkg body error")
		return
	}
	return mes,nil
}

func WritePkg(conn net.Conn,data []byte) (err error){
	//7.到这个时候　这个就是我们要发送的数据
	//7.1宪法data的长度发送给服务器
	//conn.Write(len(data))
	//先获取到data的长度->转成一个表示长度的byte切片

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	//发送长度
	n,err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(len) fail",err)
		return
	}
	//发送主体内容
	n,err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(message) fail",err)
	}
	return nil
}
