package jwtHelper

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
	conf "go-project/configs"
	"go-project/web/models/dtos"
	"strings"
	"time"
)

type TokenModel struct {
	Id string	`json:"id"`
	NickName string	`json:"nickname"`
	jwt.StandardClaims
}
//颁发JWT字符串
func IssueToken(m dtos.UserInfo) (string,error) {
	h,_:=time.ParseDuration(conf.GetConf().JwtAuth.ExpireHour)
	expireTime:=time.Now().Add(h).Unix()//时间加减后转成unix时间戳
	claims:=TokenModel{
		Id:             m.UserId,
		NickName:       m.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expireTime,//过期时间
			Issuer:conf.GetConf().JwtAuth.Issuer,//颁发机构
			//Id:"",//jti
			Audience:conf.GetConf().JwtAuth.Audience,//颁发给谁
			//Subject:"",
		},
	}
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err:=tokenClaims.SignedString([]byte(conf.GetConf().JwtAuth.SecretKey))
	return token,err
}

//解析JWT
func ParseToken(tokenStr string) (m TokenModel,err error)  {
	parseToken,err:=jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,fmt.Errorf("error x")
		}
		//自己加密的秘钥或者说盐值
		return []byte(conf.GetConf().JwtAuth.SecretKey), nil
	})
	if parseToken == nil{
		return m,err	//token格式不正确
	}
	if claims,ok:=parseToken.Claims.(jwt.MapClaims);ok && parseToken.Valid{
		m.Id=cast.ToString(claims["id"])
		m.ExpiresAt=cast.ToInt64(claims["exp"])
		return m,nil
	}else {
		return m,err
	}
}

//header中取Token
func GetToken(ctx iris.Context) (string)  {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "" // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return ""
	}

	return authHeaderParts[1]
}

//Token获取UserId
func GetUserId(token string) string  {
	userId:=""
	if token !="" && token!="undefined" && len(token)>7{
		m,_:=ParseToken(token)
		if m.Id!=""{
			userId=m.Id
		}
	}
	return userId
}