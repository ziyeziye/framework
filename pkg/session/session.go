package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func CookieStorage(name, secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	return sessions.Sessions(name, store)
}

type RedisCfg struct {
	Size                               int
	Network, Address, Password, Secret string
}

func (r RedisCfg) Default() RedisCfg {
	return RedisCfg{
		Size:     10,
		Network:  "tcp",
		Address:  "localhost:6379",
		Password: "",
		Secret:   "secret",
	}
}

func RedisStorage(name string, cfg RedisCfg) gin.HandlerFunc {
	//config = 10, "tcp", "localhost:6379", "", []byte("secret")
	store, _ := redis.NewStore(
		cfg.Size,
		cfg.Network,
		cfg.Address,
		cfg.Password,
		[]byte(cfg.Secret),
	)
	return sessions.Sessions(name, store)
}
