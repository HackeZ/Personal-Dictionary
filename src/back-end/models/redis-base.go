package models

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	// RedisClient a available Pool of Redis.
	RedisClient *redis.Pool
	// RedisDB The number of Database.
	RedisDB = 0
)

// initRedis Connect RedisDB and Return a available Pool.
// @param host string
// @return *redis.Pool
// @retrun error
func initRedis(host string) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:     64,
		IdleTimeout: 3 * time.Second,
		MaxActive:   999999, // max number of connections
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}, nil
}

// GetRedisClient return an available Redis Client.
/* Usage:
 *      GetRedisClient()
 *      // Get a Conn from Pool
 *      rc := RedisClient.Get()
 *      // Return Conn into Pool When you are Done.
 *      defer rc.Close()
 */
func GetRedisClient() {
	RedisClient, _ = initRedis(":6379")

}
