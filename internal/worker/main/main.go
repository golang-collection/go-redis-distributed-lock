package main

import (
	"go-redis-distributed-lock/global"
	"go-redis-distributed-lock/internal/worker"
	"go-redis-distributed-lock/pkg/cache"
	"go-redis-distributed-lock/pkg/setting"
	"log"
	"strings"
	"sync"
)

/**
* @Author: super
* @Date: 2021-03-18 20:36
* @Description:
**/

func Init(config string) {
	//初始化配置
	err := setupSetting(config)
	if err != nil {
		log.Printf("init setupSetting err: %v\n", err)
	} else {
		log.Printf("初始化配置信息成功")
	}
	//初始化redis
	err = setupCacheEngine()
	if err != nil {
		log.Printf("init setupCacheEngine err: %v\n", err)
	} else {
		log.Printf("初始化cache成功")
	}
}

func setupSetting(config string) error {
	newSetting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Cache", &global.CacheSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化redis pool
func setupCacheEngine() error {
	var err error
	global.RedisEngine, err = cache.NewRedisEngine(global.CacheSetting)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	Init("configs/")
	//模拟分布式的抢占
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker.Executor()
		}()
	}
	wg.Wait()
}
