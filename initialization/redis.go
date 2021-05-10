package initialization

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func Redis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", viper.Get("redis.host"), viper.Get("redis.port")),
		Password: fmt.Sprintf("%v", viper.Get("redis.password")),
		DB:       cast.ToInt(viper.Get("redis.db")),
	})
}
