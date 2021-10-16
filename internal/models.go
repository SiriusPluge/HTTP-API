package internal

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	Name string `json:"name"`
	LastName string `json:"last_name"`
	Phone int `json:"phone"`
}
