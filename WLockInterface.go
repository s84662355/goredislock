package goredislock

 

type WLockInterface  interface {
    Lock(key string , expire int ,waittime int) bool
    UnLock(key string) bool
}

 