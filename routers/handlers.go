package routers

import "github.com/gin-gonic/gin"

func userGroupHandler(r *gin.Engine) {
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
