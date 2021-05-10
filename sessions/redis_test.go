package sessions

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestRedisSessionHandler_All(t *testing.T) {
	cases := []struct {
		input map[string]string
		wants []string
	}{
		{
			input: map[string]string{"test_key": "test_value"},
			wants: []string{"test_value"},
		},
	}

	r := &RedisSessionHandler{
		Client: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
		Context: context.Background(),
	}

	i := 0

	for _, item := range cases {
		for key, value := range item.input {
			r.Set(key, value, time.Second*30)
			if r.Get(key) != item.wants[i] {
				t.Fatalf("%v != %v", r.Get(key), item.wants[i])
			}

			r.Destroy(key)

			if r.Has(key) {
				t.Fatalf("delete %v failed, it's still there.", key)
			}

			i++
		}
	}
}

func TestRedisSessionHandler_Set(t *testing.T) {
	cases := []struct {
		input map[string]string
		wants []string
	}{
		{
			input: map[string]string{"test_key": "test_value"},
			wants: []string{"test_value"},
		},
	}

	r := &RedisSessionHandler{
		Client: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
		Context: context.Background(),
	}

	i := 0

	for _, item := range cases {
		for key, value := range item.input {
			r.Set(key, value, time.Second*2)
			if r.Get(key) != item.wants[i] {
				t.Fatalf("%v != %v", r.Get(key), item.wants[i])
			}

			time.Sleep(time.Second * 3)

			if r.Has(key) {
				t.Fatalf("set %v expiry failed, it's still there.", key)
			}

			i++
		}
	}
}
