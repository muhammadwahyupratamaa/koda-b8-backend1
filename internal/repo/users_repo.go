package repo

import (
	"koda-b8-backend1/internal/model"
)

type UserRepo struct {
	data []model.User
}

func NewUserRepo() *UserRepo{
	return &UserRepo{
		data: []model.User{},
	}
}

func (r *UserRepo) Create(req *model.CreateUser) {
	id := len(r.data) + 1

	r.data = append(r.data, model.User{
		ID: int64(id),
		Email: req.Email,
		Password: req.Password,
	})
}

func (r *UserRepo) FindAll() []model.User{
	return  r.data
}

func (r *UserRepo) findByEmail(email string) *model.User {
	for _, user := range r.data{
		if user.Email == email {
			return &user
		}
	}
	return nil
}
