package sessions

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type MysqlSessionHandler struct {
	Client *gorm.DB
}

var _ SessionHandler = &MysqlSessionHandler{}

func (m *MysqlSessionHandler) Get(key string) string {
	session := &SessionModel{}
	result := m.Client.Where("`key` = ?", key).First(session)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ""
	}

	if time.Now().After(session.UpdatedAt.Add(session.Expiry)) {
		return ""
	}

	return session.Value
}

func (m *MysqlSessionHandler) Set(key, value string, duration time.Duration) {
	err := m.Client.Transaction(func(tx *gorm.DB) error {
		session := &SessionModel{}
		result := tx.Where("`key` = ?", key).First(session)

		session.Key = key
		session.Value = value
		session.Expiry = duration

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if err := tx.Create(session).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Save(session).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		panic(err.Error())
	}
}

func (m *MysqlSessionHandler) Has(key string) bool {
	session := &SessionModel{}
	result := m.Client.Where("`key` = ?", key).First(session)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func (m *MysqlSessionHandler) Destroy(key string) bool {
	result := m.Client.Where("`key` = ?", key).Unscoped().Delete(&SessionModel{})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func (m *MysqlSessionHandler) Migrate() error {
	return m.Client.AutoMigrate(&SessionModel{})
}
