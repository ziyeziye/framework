package cache

import (
	"fmt"
	"framework/config"
	"gopkg.in/redis.v5"
)

var globalRedis *redis.Client

type RedisDb struct {
	Conn *redis.Client
}

func factory() *redis.Client {
	conf := config.GetSec("redis")
	host := conf.Key("HOST").String()
	port := conf.Key("PORT").String()
	password := conf.Key("PASS").MustString("")
	fmt.Printf("conf-redis: %s:%s - %s\r\n", host, port, password)

	address := fmt.Sprintf("%s:%s", host, port)
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
		PoolSize: conf.Key("POOL_SIZE").MustInt(5),
	})
}

func init() {
	globalRedis = factory()
	NewRedis().RedisCheck()
}

func NewRedis() *RedisDb {
	return RedisDb{}.GetConn()
}

func (r RedisDb) GetConn() *RedisDb {
	r.Conn = globalRedis
	return &r
}

/**
 * 获取master连接
 */
func RedisMaster() *redis.Client {
	return RedisDb{}.GetConn().Conn
}

/**
 * 检查连接是否正常
 */
func (r *RedisDb) RedisCheck() {
	_, err := r.Conn.Ping().Result()
	if err != nil {
		panic(err)
	}
}
