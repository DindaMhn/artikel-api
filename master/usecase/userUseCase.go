package usecase

import "artikel/master/model"

type UserUseCase interface {
	GetAllDataUser() ([]*model.UserModel, error)
	GetUserById(id string) (model.UserModel, error)
	UpdateUserById(id string, DataUser model.UserModel) error
	DeleteUserById(id string) error
	InsertUser(DataUser model.UserModel) error
}
