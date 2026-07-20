package model

import "time"


type User struct {
	ID int64
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUser struct{
	Email string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type LoginUser struct {
	Email string `form:"email"`
	Password string `form:"password"`
}

