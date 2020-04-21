package goredislock

import (
	"goredislock/script"
	"time"
	"fmt"
	"github.com/go-redis/redis"
)

type WLock struct {
	clientname string
    redisClient * redis.Client
}

func (l *WLock) Lock(key string , expire int ,waittime int) bool{

	for waittime >=0 {

		res:=l.lockEval(key , expire)  

		fmt.Println("res:: ",res , "    " , waittime )

		if res == true {
			return true
		}

		time.Sleep(time.Duration(1)*time.Millisecond )
		waittime--
	}

	return  false
}

func (l *WLock) UnLock(key string) bool{
	res := l. unlockEval(key) 
	if res == true{
		return  true
	}
	return  false
}

func (l *WLock)  lockEval(key string , expire int ) bool{
	var updateRecordExpireScript = redis.NewScript(script.GetWLock())
	res , err := updateRecordExpireScript.Run(l.redisClient, []string{ key },expire,l.clientname ).Result()
	if err != nil {
		variant := fmt.Sprintf(" 脚本执行失败 %s ", err)
		panic(variant)
	}
	if  res.(int64) == 1{
		return  true
	}
	return  false
}

func (l *WLock)  unlockEval(key string ) bool{
	var updateRecordExpireScript = redis.NewScript(script.GetUnWLock())
	res , err := updateRecordExpireScript.Run(l.redisClient,  []string{ key }, l.clientname).Result()
	if err != nil {
		variant := fmt.Sprintf(" 脚本执行失败 %s ", err)
		panic(variant)
	}
	if  res.(int64) == 1{
		return  true
	}
	return  false

}


func NewWLock(redisClient * redis.Client,clientname string) *WLock{
       client := new(WLock)
       client.redisClient = redisClient
       client.clientname = clientname
       return  client
}
