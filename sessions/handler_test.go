package sessions

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestSession_All(t *testing.T) {
	cases := []struct {
		input map[string]string
		wants []string
	}{
		{
			input: map[string]string{"test_key": "test_value"},
			wants: []string{"test_value"},
		},
	}

	session := New(&RedisSessionHandler{
		Client: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
		Context: context.Background(),
	}, time.Hour)

	i := 0

	for _, item := range cases {
		for key, value := range item.input {
			session.Set(key, value)
			if session.Get(key) != item.wants[i] {
				t.Fatalf("%v != %v", session.Get(key), item.wants[i])
			}

			session.Destroy(key)

			if session.Has(key) {
				t.Fatalf("delete %v failed, it's still there.", key)
			}

			i++
		}
	}
}
