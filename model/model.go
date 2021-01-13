package model

import "encoding/json"

type ToString interface {
	String() string
}

type StringableModel struct {
}

func (m *StringableModel) String() string {
	res, _ := json.Marshal(m)
	return string(res)
}
