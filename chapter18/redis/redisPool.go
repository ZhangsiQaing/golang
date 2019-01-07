package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

//定义一个全局的pool
var pool *redis.Pool


//当启动程序时,就初始化连接池
func init(){
	pool = &redis.Pool{
		MaxIdle:8,   //最大空闲连接数
		MaxActive:0,  //最大连接数，０表示没有限制
		IdleTimeout:100,  //最大空闲时间
		Dial:func() (redis.Conn,error) {//初始化连接代码，连接那个ip
			conn,err := redis.Dial("tcp","172.16.1.14:6379")
			conn.Do("AUTH","AjYy366JVr@ju@K9")
			return conn,err
		},
	}
}

func main(){
	//先从pool　取出一个连接
	//pool.Close()
	conn := pool.Get()
	defer conn.Close()


	res,err := conn.Do("HSET","student","name","jac")
	fmt.Println(res,err)
	res1,err := redis.String(conn.Do("HGET","student","name"))
	fmt.Printf("res:%s,error:%v",res1,err)

}