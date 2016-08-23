package models

import (
	"log"
	"testing"
)

func TestPing(t *testing.T) {
	rc := RedisClient.Get()
	defer rc.Close()

	if pong, err := rc.Do("PING"); pong != "PONG" || err != nil {
		log.Println(err)
		t.Error("Can not PING Redis Server.")
	}
}
