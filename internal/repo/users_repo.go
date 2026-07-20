package repo

import (
	"koda-b8-backend1/internal/model"
)

type UserRepo struct {
	data *[]model.User
}

func NewUserRepo(data *[]model.User) *UserRepo {
	return &UserRepo{
		data: data,
	}
}

func (r *UserRepo) Create(req *model.CreateUser) {

	id := len(*r.data) + 1

	*r.data = append(*r.data, model.User{
		ID:       int64(id),
		Email:    req.Email,
		Password: req.Password,
	})
}

func (r *UserRepo) FindAll() []model.User {
	return *r.data
}

func (r *UserRepo) FindByEmail(email string) *model.User {

	for i := range *r.data {

		if (*r.data)[i].Email == email {
			return &(*r.data)[i]
		}

	}

	return nil
}