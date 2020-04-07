package router
import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-project/web/controllers"
	"go-project/web/middleware"
)

func Register(app *iris.Application){
	baseUrl:="/api"//路由前缀

	//system控制器
	sysParty:=app.Party(baseUrl+"/sys")
	mvc.New(sysParty).Handle(controllers.NewSysController())

	//用户控制器
	userParty:=app.Party(baseUrl+"/user")
	userParty.Use(middleware.JwtAuthorization)//需要授权中间件
	mvc.New(userParty).Handle(controllers.NewUserController())
}