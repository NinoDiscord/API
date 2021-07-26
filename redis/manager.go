package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"strings"
	"time"
)

type Redis struct {
	Connection *redis.Client
}

func NewRedisClient() *Redis {
	return &Redis{
		Connection: nil,
	}
}

func (r *Redis) Connect() error {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB")); if err != nil {
		return err
	}

	password := os.Getenv("REDIS_PASSWORD")
	if sentinels := os.Getenv("REDIS_SENTINELS"); len(sentinels) > 0 {
		hosts := strings.Split(os.Getenv("REDIS_SENTINELS"), ";")
		r.Connection = redis.NewFailoverClient(&redis.FailoverOptions{
			SentinelAddrs: hosts,
			MasterName: os.Getenv("REDIS_MASTER"),
			Password: password,
			DB: db,
			DialTimeout: 10 * time.Second,
			ReadTimeout: 15 * time.Second,
			WriteTimeout: 15 * time.Second,
		})
	} else {
		host := os.Getenv("REDIS_HOST")
		port := os.Getenv("REDIS_PORT")
		r.Connection = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB: db,
			DialTimeout: 10 * time.Second,
			ReadTimeout: 15 * time.Second,
			WriteTimeout: 15 * time.Second,
		})
	}

	// check if connection is healthy
	if err := r.Connection.Ping(context.TODO()).Err(); err != nil {
		return err
	}

	return nil
}
