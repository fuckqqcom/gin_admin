package routers

import (
	"gin_admin/pkg/engine"
	"github.com/gin-gonic/gin"
)

func InitRouter() engine.GEng {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//加载userGroupHandler路由
	userGroupHandler(r)
	return r
}
