package main

import(
	"goredislock"
	"github.com/go-redis/redis"
	"fmt"
	//"strconv"
	)

var client = redis.NewClient(&redis.Options{Addr:"127.0.0.1:6379",Password: "", DB:0})

 

 

func main(){
	lock := goredislock.NewWLock( client , "asds2a222sda")
	if( lock.Lock("1111a", 20 ,1000) ){
		fmt.Println("bbbbbbbbb")

		fmt.Println( lock.UnLock("1111a") )
		return
	}
	fmt.Println("ffffffff")
}