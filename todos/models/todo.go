package models

import (
	configs "github.com/crowdeco/bima/configs"
)

type Todo struct {
	configs.Base

    Name string

}

func (m *Todo) TableName() string {
	return "todo"
}

func (m *Todo) IsSoftDelete() bool {
	return true
}
