package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string
	Password string
}
