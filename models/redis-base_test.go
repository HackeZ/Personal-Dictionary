package models

import (
	"fmt"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestConn(t *testing.T) {
	fmt.Println("=== TestConn Start ===")

	GetRedisClient()
	rc := RedisClient.Get()
	defer rc.Close()

	if pong, err := rc.Do("PING"); pong != "PONG" || err != nil {
		t.Error("Can not PING Redis Server.")
	}

	if _, err := rc.Do("SET", "ACP_TEST1", "test"); err != nil {
		t.Error(err.Error())
	}

	if testValue, _ := redis.String(rc.Do("GET", "ACP_TEST1")); testValue != "test" {
		t.Error("GET KEY ACP_TEST1 VALUE WRONG!")
	}

	if _, err := rc.Do("DEL", "ACP_TEST1"); err != nil {
		t.Error(err.Error())
	}

	fmt.Println("=== TestConn End ===")
}
