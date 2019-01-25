package engine

import (
	"../../pkg/err"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Send(httpCode, code int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": code,
		"msg":  err.GetMsg(code),
		"data": data,
	})
}
