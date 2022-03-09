package repo

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RedisRepository) Set(key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(key, value, ttl).Err()
}

func (r *RedisRepository) LPush(key string, value ...interface{}) (int64, error) {
	return r.client.LPush(key, value...).Result()
}

func (r *RedisRepository) RPush(key string, value ...interface{}) (int64, error) {
	return r.client.RPush(key, value...).Result()
}

func (r *RedisRepository) LPop(key string) (string, error) {
	return r.client.LPop(key).Result()
}

func (r *RedisRepository) RPop(key string) (string, error) {
	return r.client.RPop(key).Result()
}

func (r *RedisRepository) IfExist(key string) (bool, error) {
	val, err := r.client.Exists(key).Result()
	if err != nil {
		return false, err
	} else {
		return val == 1, nil
	}
}

func (r *RedisRepository) BRPop(timeout time.Duration, key string) ([]string, error) {
	return r.client.BRPop(timeout, key).Result()
}

func (r *RedisRepository) Llen(key string) int64{
	return r.client.LLen(key).Val()
}