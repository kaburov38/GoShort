package model

import (
	"gorm.io/gorm"
)

type Mapping struct {
	gorm.Model
	Source string
	Target string
}

func (Mapping) TableName() string {
	return "mappings"
}
