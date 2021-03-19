package global

import "github.com/garyburd/redigo/redis"

/**
* @Author: super
* @Date: 2021-03-18 19:59
* @Description:
**/

var (
	RedisEngine *redis.Pool
)

func GetConn() redis.Conn {
	return RedisEngine.Get()
}
