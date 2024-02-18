package users

import "time"

type User struct {
	UserName  interface{}
	LastName  string
	FirstName string
	Roles     map[string]interface{}
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type DbUser struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Email     *string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
