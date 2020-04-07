package service

import (
	"go-project/utils/jwtHelper"
	"go-project/web/models"
	"go-project/web/models/dtos"
)

type SysService interface {
	Login(m map[string]string) (res models.Result)
}

type sysService struct {}

func NewSysService() SysService  {
	return &sysService{}
}


func (u sysService) Login(m map[string]string) (res models.Result)  {
	user:=dtos.UserInfo{Username:"hzh",UserId:"111",Token:""}
	token,_ :=jwtHelper.IssueToken(user)
	user.Token=token
	res.Code=1
	res.Data=user
	res.Msg="success"
	return
}