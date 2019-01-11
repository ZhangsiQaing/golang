package utils

import (
	"net"
	"syst/chatroom/common/message"
	"fmt"
	"encoding/binary"
	"encoding/json"
)

type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf [8096]byte //这是传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message,err error){
	//buf := make([]byte,8096)
	fmt.Println("读取客户端发送的数据....")
	_,err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}

	//根据buf[:4] 转成一个　uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据 pkgLen 读取消息内容
	n,err := this.Conn.Read(this.Buf[:pkgLen])
	//fmt.Println(n,string(buf),err)
	if n != int(pkgLen) || err != nil {
		//fmt.Println("connn.Read fail err = ",err)
		//err = errors.New("read pkg header error")
		return
	}

	//buf反序列化成 -> message.Message
	err = json.Unmarshal(this.Buf[:n],&mes)
	fmt.Println(err)
	if err != nil {
		//fmt.Println("json.Unmarshal fail err = ",err)
		//err = errors.New("read pkg body error")
		return
	}
	return mes,nil
}

func (this *Transfer) WritePkg(data []byte) (err error){
	//7.到这个时候　这个就是我们要发送的数据
	//7.1宪法data的长度发送给服务器
	//conn.Write(len(data))
	//先获取到data的长度->转成一个表示长度的byte切片

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)
	//发送长度
	n,err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(len) fail",err)
		return
	}
	//发送主体内容
	n,err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(message) fail",err)
	}
	return nil
}
