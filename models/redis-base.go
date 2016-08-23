package models

import (
	"log"
	"time"

	"github.com/astaxie/beego"

	"github.com/garyburd/redigo/redis"
)

var (
	// RedisClient a available Redis Client.
	RedisClient *redis.Pool
	// REDIS_DB The number of DB in Redis.
	REDIS_DB int
	// REDIS_HOST Host
	REDIS_HOST string
)

func init() {
	REDIS_HOST = beego.AppConfig.String("redis_host")
	REDIS_DB, _ = beego.AppConfig.Int("redis_db")

	if REDIS_HOST == "" {
		REDIS_HOST = ":6379"
	}
	if REDIS_DB == 0 {
		log.Println("May Your Setting Have Not Set 'redis_db', use default ==> 0")
	}

	RedisClient = &redis.Pool{
		MaxIdle:     64,
		IdleTimeout: 3 * time.Second,
		MaxActive:   99999,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// Choose DB
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}
