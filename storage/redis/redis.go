package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/Asliddin3/exam-api-gateway/storage/repo"
	"github.com/gomodule/redigo/redis"
)

type redisRepo struct {
	rds *redis.Pool
}

func NewRedisRepo(rds *redis.Pool) repo.RedisRepo {
	return &redisRepo{
		rds: rds,
	}
}
func (r *redisRepo) Set(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(3))
	defer cancel()
	conn := r.rds.Get()
	_, err := conn.Do("SET", ctx, key, value, 0)
	if err != nil {
		return err
	}

	return nil
}

func (r *redisRepo) SetWithTTL(key, value string, seconds int) error {
	conn := r.rds.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, seconds, value)
	return err
}

func (r *redisRepo) Get(key string) (interface{}, error) {
	fmt.Println(key)
	conn := r.rds.Get()
	defer conn.Close()
	return conn.Do("GET", key)
}
