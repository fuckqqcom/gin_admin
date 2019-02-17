package api

import (
	"gin_admin/pkg/engine"
	"gin_admin/pkg/err"
	"gin_admin/tools/utils"
	"github.com/gin-gonic/gin"
)

type LUser struct {
	UserAccount string `json:"user_account"  binding:"required"`
	Pwd         string `json:"pwd"  binding:"required"`
	GroupId     string `json:"group_id"  binding:"required"`
}

func Register(c *gin.Context) {
	var lu LUser
	ctx := engine.Gin{c}
	if !utils.CheckError(c.BindJSON(&lu), "Register") {
		ctx.Send(200, err.Error, "")
		return
	}
	/**
	调用微服务
	*/
}

func Login(c *gin.Context) {
	var lu LUser
	ctx := engine.Gin{c}
	if !utils.CheckError(c.BindJSON(&lu), "Register") {
		ctx.Send(200, err.Error, "")
		return
	}
}
