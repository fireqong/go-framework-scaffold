package sessions

import (
	"time"

	"gorm.io/gorm"
)

type SessionModel struct {
	gorm.Model

	Key    string        `gorm:"type:varchar(128);uniqueIndex"`
	Value  string        `gorm:"type:longtext"`
	Expiry time.Duration `gorm:"type:int"`
}

func (SessionModel) TableName() string {
	return "sessions"
}
