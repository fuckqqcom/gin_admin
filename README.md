# gin_admin

[![Build Status](https://travis-ci.org/go-sh/gin_admin.svg?branch=master)](https://travis-ci.org/go-sh/gin_admin)
[![GoDoc](https://godoc.org/github.com/go-sh/gin_admin?status.svg)](https://godoc.org/github.com/go-sh/gin_admin)
[![Open Source Helpers](https://www.codetriage.com/go-sh/gin_admin/badges/users.svg)](https://www.codetriage.com/go-sh/gin_admin)

## TODO

项目重构（2019.2.17）


## 工程目录说明

- `api` 对外提供微服务
- `service` 内部通信
- `pkg` 放一些公共的文件,比如加载配置文件,生成和解析token 自定义错误状态码等
- `config` 放配置文件

## 技术选型
- `Go-micro` 主要使用的微服务框架
- `xorm/gorm` 选用 ORM
- `Gin` 做路由蹭
- `Consul` 做服务注册和发现, `etcd`也行(后面也可插件)

## 服务调用流程：

Api---gin(路由) -> 对应 handler -> 对应逻辑层 -> 调用相应微服务接口

## 服务拆分


    微服务:user:
		用户
		    Uid string //用户id
		    Account string //账号
		    Phone string //手机号码
		    Addr string //地址
		    Email string //邮箱
		    IsValid int //是否禁用 -1表示禁用 1表示正常使用
		    IsDel int //是否删除 //逻辑删除 -1表示删除 1表示正常用户


		角色
		    Name string //角色名称
		    RoleId string //角色id
		    IsValid int //是否禁用 -1表示禁用 1表示正常使用
            IsDel int //是否删除 //逻辑删除 -1表示删除 1表示正常

		权限
		    Pid string //权限id
		    Method string
		    Url string //url
		    这里只是针对api还是针对所有的包括目录什么的
		    ParentId string // 0表示父id,非0表示子id
		    IsValid int //是否禁用 -1表示禁用 1表示正常使用
            IsDel int //是否删除 //逻辑删除 -1表示删除 1表示正常
		组
            Name string //组名称
            GroupId string //组id
            IsValid int //是否禁用 -1表示禁用 1表示正常使用
            IsDel int //是否删除 //逻辑删除 -1表示删除 1表示正常

	就是组->用户->角色->权限这样的流程


然后在gin中做一个middleware,通过角色统一做api鉴权等

可以做成通用的用户权限认证管理,后面其他想扩展的做成对应的微服务，比如做一个blog，直接一个blog微服务就好
