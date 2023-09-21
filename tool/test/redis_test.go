package test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.146.146:6379",
		Password: "2001823",
		DB: 0,
		PoolSize: 20,
	})

	if s, err := client.Ping(context.Background()).Result(); err != nil {
		t.Fatal(err)
	} else {
		
		t.Log(s)
	}
}


func TestRedisSet(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.146.146:6379",
		Password: "2001823",
		DB: 0,
		PoolSize: 20,
	})

	if _, err := client.Set(context.Background(), "k1", "v1", time.Second * 30).Result(); err != nil {
		t.Fatal(err)
	} 
}

func TestRedisSetnGet(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.146.146:6379",
		Password: "2001823",
		DB: 0,
		PoolSize: 20,
	})

	if _, err := client.Set(context.Background(), "k1", "v1", time.Second * 30).Result(); err != nil {
		t.Fatal(err)
	} 

	if s, err := client.Get(context.Background(), "k1").Result(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(s)
	}
}