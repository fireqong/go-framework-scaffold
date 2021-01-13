package sessions

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestMysqlSessionHandler_Migrate(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:fuckyou123@tcp(127.0.0.1:3306)/go?parseTime=true&charset=utf8mb4"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	m := &MysqlSessionHandler{
		Client: db,
	}

	if err := m.Migrate(); err != nil {
		panic(err.Error())
	}
}

func TestMysqlSessionHandler_All(t *testing.T) {
	cases := []struct {
		input map[string]string
		wants []string
	}{
		{
			input: map[string]string{"test_key": "test_value"},
			wants: []string{"test_value"},
		},
	}

	db, err := gorm.Open(mysql.Open("root:fuckyou123@tcp(127.0.0.1:3306)/go?parseTime=true&charset=utf8mb4"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	r := &MysqlSessionHandler{
		Client: db,
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
