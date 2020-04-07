package middleware

import (
	"github.com/kataras/iris/v12"
	"go-project/utils/jwtHelper"
	"go-project/web/models"
)

//检测Header请求头验证Token
func JwtAuthorization(ctx iris.Context)  {
	token:=jwtHelper.GetToken(ctx)
	//token不为空
	if token ==""{
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(models.Result{Code:-1,Msg:"token param error!"})//token需添加 bearer 前缀
		return
	}
	m,err:=jwtHelper.ParseToken(token)
	if err!=nil{
		//logger
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(models.Result{Code:-1,Msg:"token is expired!"})
		return
	}
	if m.Id==""{
		//token无效
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(models.Result{Code:-1,Msg:"token格式错误"})
		return
	}

	ctx.Next()
}
