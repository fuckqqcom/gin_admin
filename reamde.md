api 对外提供挨批服务
service 内部通信
pkg 放一些公共的文件,比如加载配置文件,生成和解析token 自定义错误状态码等
config 放配置文件




技术选型Go-micro  gin  xorm/gorm
Gin做路由蹭
Consul做服务注册和发现/etcd也行(后面也可插件)

Api---gin(路由)—对应的函数-->微服务

微服务”user:
		用户
		角色
		权限
		组
	就是组->用户->角色->权限这样的流程

然后在gin中做一个middleware,通过角色统一做api鉴权等

可以做成通用的用户权限认证管理,后面其他想扩展的做成对应的微服务，比如做一个blog，直接一个blog微服务就好
