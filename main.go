package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
)

var Conn redis.Conn

func init() {
	Network := "tcp"
	Address := "127.0.0.1:46379"

	//连接数据库
	database, err := redis.Dial(Network, Address)
	if err != nil {
		panic(err)
	}

	//认证数据库
	database.Do("auth", "password")

	Conn = database
}

//Redis 字符串(String)
func stringFunc() {
	//写入
	_, err := Conn.Do("set", "gon", 44)
	if err != nil {
		panic(err)
	}

	//读取，针对字符串，一定要加redis.String
	re, err := redis.String(Conn.Do("get", "gon"))
	if err != nil {
		panic(err)
	}

	fmt.Println(re)
	fmt.Printf("%T", re)
}

//hash
func hashFunc() {

	hashdb := "myhashdb"

	rand.Seed(time.Now().UnixNano())
	hval := string(fmt.Sprintf("%05v", rand.Intn(1000)))
	hkey := "name" + hval
	_, err := Conn.Do("hset", hashdb, hkey, hval)
	if err != nil {
		panic(err)
	}

	re, err := redis.Values(Conn.Do("hgetall", hashdb))
	if err != nil {
		panic(err)
	}
	hm := make(map[string]string)
	for i := 0; i < len(re); i += 2 {
		hm[fmt.Sprintf("%s", re[i])] = fmt.Sprintf("%s", re[i+1])
	}

	fmt.Println(hm)
}

//List
func listFunc() {
	lk := "runoobkey"
	lv := "x"

	//lpush 数据
	_,err := Conn.Do("lpush",lk,lv)
	if err != nil {
		panic(err)
	}

	//lrange
	values,err := redis.Values(Conn.Do("lrange",lk,0,-1))
	if err != nil {
		panic(err)
	}

	for _,v := range values {
		fmt.Println(string(v.([]byte)))
	}

	//lpop 数据
	re,err := redis.String(Conn.Do("lpop",lk))
	if err != nil {
		panic(err)
	}

	fmt.Println(re)


	//redis llen
	res,err := Conn.Do("llen",lk)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)


}

func main() {
	//Redis 字符串(String) 操作
	stringFunc()

	//Redis Hash 操作
	hashFunc()

	//Redis List 操作
	listFunc()
}
