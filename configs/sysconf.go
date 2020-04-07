package configs

import (
	"github.com/json-iterator/go"
	"io/ioutil"
)

var sysConf configInfo

func GetConf() configInfo{
	return sysConf
}

func init()  {
	content,err :=ioutil.ReadFile("./configs/appsetting.json")
	if err!=nil{
		panic("sysconfig read error")
	}
	err = jsoniter.Unmarshal(content,&sysConf)//A high-performance 100% compatible drop-in replacement of "encoding/json"
	if err!=nil{
		panic(err)
	}
}

type configInfo struct {
	Dbconn string
	Server serverDto
	JwtAuth JwtAuthDto
}
type serverDto struct {
	Host string
	Port string
}
type JwtAuthDto struct {
	SecretKey string //签名密钥
	Issuer string//颁发机构
	Audience string//颁发给谁
	ExpireHour string//过期时间，单位小时
}