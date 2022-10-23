package redis

import (
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

func (r *redisRepo) SetWithTTL(key, value string, seconds int) error {
	conn := r.rds.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, seconds, value)
	return err
}

func (r *redisRepo) Get(key string) (interface{}, error) {
	conn := r.rds.Get()
	defer conn.Close()
	return conn.Do("GET", key)
}
