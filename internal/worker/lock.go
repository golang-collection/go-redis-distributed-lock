package worker

import (
	"context"
	"errors"
	"github.com/garyburd/redigo/redis"
	"go-redis-distributed-lock/global"
	"log"
	"time"
)

/**
* @Author: super
* @Date: 2021-03-18 20:02
* @Description: redis实现分布式锁
**/
type RedisLock struct {
	Key        string
	TTL        int64 //锁超时时间
	IsLocked   bool
	CancelFunc context.CancelFunc //用于结束当前执行任务，即自动续约终止
}

func NewRedisLock(key string) *RedisLock {
	redisLock := &RedisLock{
		Key: key,
		TTL: 3,
	}
	return redisLock
}

func (r *RedisLock) TryLock() (err error) {
	//申请租约上锁
	if err = r.Grant(); err != nil {
		return
	}
	ctx, cancelFunc := context.WithCancel(context.TODO())
	r.CancelFunc = cancelFunc
	//续约
	r.KeepAlive(ctx)
	r.IsLocked = true
	return nil
}

func (r *RedisLock) UnLock() (err error) {
	var res int
	if r.IsLocked {
		if res, err = redis.Int(global.GetConn().Do("DEL", r.Key)); err != nil {
			return errors.New("释放锁失败")
		}
		if res == 1 {
			//释放锁成功，调用cancelFunc
			r.CancelFunc()
			return
		}
	}
	return errors.New("释放锁失败")
}

func (r *RedisLock) KeepAlive(ctx context.Context) {
	// 类似于ETCD里面续约实现启用单独携程进行操作
	go func() {
		for {
			select {
			// 已经调用cancelFunc
			case <-ctx.Done():
				return
			default:
				// 自动续约，延长当前key结束时间
				res, err := redis.Int(global.GetConn().Do("EXPIRE", r.Key, r.TTL))
				if err != nil {
					log.Println("自动续约释放", err)
				}
				if res != -1 {
					log.Println("自动续约")
				}
				time.Sleep(time.Duration(r.TTL/2) * time.Second)
			}
		}
	}()
}

// 上锁
func (r *RedisLock) Grant() (err error) {
	if res, err := redis.String(global.GetConn().Do("SET", r.Key, "1", "NX", "EX", r.TTL)); err == nil {
		if res == "OK" {
			return nil
		}
	}
	return errors.New("上锁失败")
}
