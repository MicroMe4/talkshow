package redis

import "github.com/go-redis/redis/v7"

//GetRedisClient 获取redis客户端
func GetRedisClient(opt *redis.Options) *redis.Client{
	return redis.NewClient(opt)
}