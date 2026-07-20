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
	Email string `json :"email"`
	Password string `json:"-"`
}


