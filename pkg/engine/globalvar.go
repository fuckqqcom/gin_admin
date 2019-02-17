package engine

import (
	"../../pkg/err"
	"github.com/gin-gonic/gin"
)

type G struct {
	C *gin.Context
}

//xiaohan 2019-02-17 edit 全局配置，替换路由框架的时候只需要修改这一处
type Gin = *gin.Context

type GEng = *gin.Engine

func (g *G) Send(httpCode, code int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": code,
		"msg":  err.GetMsg(code),
		"data": data,
	})
}
