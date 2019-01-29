package main

import "gin_admin/pkg/config"

func Connection() {
	path := "config/config.json"
	config.InitConfig(path)
}
