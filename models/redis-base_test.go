package models

import (
	"testing"
)

func PingTest(t *testing.T) {
	GetRedisClient()
	rc := RedisClient.Get()
	defer rc.Close()

	if pong, err := rc.Do("PING"); pong != "PONG" || err != nil {
		t.Error("Can not PING Redis Server.")
	}
}
