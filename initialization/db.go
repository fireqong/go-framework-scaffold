package initialization

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	database := cast.ToStringMapString(viper.Get("database"))
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?%v",
		database["username"],
		database["password"],
		database["host"],
		database["port"],
		database["database"],
		database["options"],
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
