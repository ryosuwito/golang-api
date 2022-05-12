package structs

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Status   bool   `json:"status" gorm:"default:true"`
}

type UsersLogin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *Users) BeforeUpdate(tx *gorm.DB) (err error) {
	status := u.Status
	if !status {
		u.Status = false
	}
	return
}

type Products struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
