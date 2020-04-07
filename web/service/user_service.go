package service

import (
	"go-project/web/models"
	"go-project/web/repositories"
)

type UserService interface {
	//Login(m map[string]string) (res models.Result)
	Add(a int)(res models.Result)
}

type userService struct {}

var userRepository =repositories.NewUserRepository()

func NewUserService() UserService  {
	return &userService{}
}

func (s userService) Add(a int)(res models.Result)  {
	flag:=	userRepository.Add("aaa")
	return models.Result{Data:flag,Msg:"ok",Code:1}
}
