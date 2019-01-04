package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)


//'redis' => [
//'hostname' => '172.16.1.14',
//'port' => 6379,
//'password' => 'AjYy366JVr@ju@K9',
//'database' => 0,
//],

func main(){
	conn,err := redis.Dial("tcp","172.16.1.14:6379")
	if err != nil {
		fmt.Println("redis.Dial err = ",err)
		return
	}
	defer conn.Close()

	conn.Do("AUTH","AjYy366JVr@ju@K9")

	//fmt.Println("conn succ...",conn)
	_,err = conn.Do("set","name","tomjerry")
	if err != nil{
		fmt.Println("redis set err = ",err)
		return
	}

	r,err := redis.String(conn.Do("get","name"))
	if err != nil{
		fmt.Println("redis get err = ",err)
		return
	}else{
		fmt.Println("操作ok name = ",r)
	}



}