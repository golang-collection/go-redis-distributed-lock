package worker

import (
	"fmt"
	"log"
	"time"
)

/**
* @Author: super
* @Date: 2021-03-18 22:10
* @Description: 执行正常的业务逻辑
**/

const LOCK = "redislock"

func Executor() {
	locker := NewRedisLock(LOCK)
	err := locker.TryLock()
	if err != nil {
		log.Println(err)
		return
	}
	defer locker.UnLock()

	//执行业务逻辑
	fmt.Println("hello")
	time.Sleep(5 * time.Second)
	fmt.Println("finish")
}
