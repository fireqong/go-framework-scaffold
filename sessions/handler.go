package sessions

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"time"
)

type SessionHandler interface {
	Get(string) string
	Set(string, string, time.Duration)
	Has(string) bool
	Destroy(string) bool
}

type Session struct {
	SessionId string
	Expiry    time.Duration
	Driver    SessionHandler
}

func New(driver SessionHandler, expiry time.Duration) *Session {
	sessionId := GenerateSessionId()
	return &Session{
		Driver:    driver,
		SessionId: sessionId,
		Expiry:    expiry,
	}
}

func NewWithSessionId(driver SessionHandler, sessionId string, expiry time.Duration) *Session {
	return &Session{
		SessionId: sessionId,
		Driver:    driver,
		Expiry:    expiry,
	}
}

func (s *Session) Get(key string) string {
	decodeValue := s.GetAll()
	return decodeValue[key]
}

func (s *Session) Set(key, value string) {
	decodeValue := s.GetAll()
	decodeValue[key] = value

	encodeValue, err := json.Marshal(decodeValue)

	if err != nil {
		panic(err.Error())
	}

	s.Driver.Set(s.SessionId, string(encodeValue), s.Expiry)
}

func (s *Session) Has(key string) bool {
	value := s.Get(key)
	if value == "" {
		return false
	}

	return true
}

func (s *Session) Destroy(key string) bool {
	if s.Has(key) {
		decodeValue := s.GetAll()
		delete(decodeValue, key)

		encodeValue, err := json.Marshal(decodeValue)

		if err != nil {
			panic(err.Error())
		}

		s.Driver.Set(s.SessionId, string(encodeValue), s.Expiry)
	}

	return true
}

func (s *Session) DestroyAll() bool {
	return s.Driver.Destroy(s.SessionId)
}

func (s *Session) GetAll() map[string]string {
	decodeValue := map[string]string{}

	encodeValue := "{}"
	if s.Driver.Has(s.SessionId) {
		encodeValue = s.Driver.Get(s.SessionId)
	}

	err := json.Unmarshal([]byte(encodeValue), &decodeValue)

	if err != nil {
		panic(err.Error())
	}

	return decodeValue
}

func GenerateSessionId() string {
	return uuid.NewV4().String()
}
