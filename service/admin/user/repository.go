package user

import (
	"gin_admin/pkg/config"
	"gin_admin/service/admin/proto"
)

type RepositoryUser interface {
	Register(*admin.UserParams) *admin.Response
	Login(**admin.UserParams) *admin.Response
}

func init() {
	config.InitConfig("config/config.json")
}

func Register(account, pwd, groupId string) {
	config.EngDb.Insert()
}
