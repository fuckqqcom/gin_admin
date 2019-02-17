package routers

import (
	"gin_admin/pkg/engine"
	"github.com/gin-gonic/gin"
)

func userGroupHandler(r engine.GEng) {
	u := r.Group("user")
	{
		u.GET("/ping", ping)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}
