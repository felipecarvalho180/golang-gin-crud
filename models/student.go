package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name"`
}

func (Student) TableName() string {
	return "students"
}
