package sessions

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisSessionHandler struct {
	Client  *redis.Client
	Context context.Context
}

var _ SessionHandler = &RedisSessionHandler{}

func (r *RedisSessionHandler) Get(key string) string {
	val, err := r.Client.Get(r.Context, key).Result()

	if err != nil {
		panic(err.Error())
	}

	return val
}

func (r *RedisSessionHandler) Set(key string, value string, expiration time.Duration) {
	setCmd := redis.NewScript(`
		local ttl
		if redis.call('exists', KEYS[1]) == 0 then
			ttl=ARGV[1]
		else
			ttl=redis.call('ttl', KEYS[1])
		end
		return redis.call('setex', KEYS[1], ttl, ARGV[2])
	`)

	err := setCmd.Run(r.Context, r.Client, []string{key}, int(expiration/time.Second), value).Err()

	if err != nil {
		panic(err.Error())
	}
}

func (r *RedisSessionHandler) Has(key string) bool {
	res, err := r.Client.Exists(r.Context, key).Result()

	if err != nil {
		panic(err.Error())
	}

	if res == 0 {
		return false
	}

	return true
}

func (r *RedisSessionHandler) Destroy(key string) bool {
	res, err := r.Client.Del(r.Context, key).Result()

	if err != nil {
		panic(err.Error())
	}

	if res == 0 {
		return false
	}

	return true
}
