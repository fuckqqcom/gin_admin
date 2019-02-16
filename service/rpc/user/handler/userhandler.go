package handler

//TODO 重构

// import (
// 	"context"
// )

// type Handler struct {
// 	Repo RepositoryUser
// }

// func (h *Handler) Register(ctx context.Context, req *admin.UserParams) (resp *admin.Response) {
// 	/**
// 	注册流程,先验证GroupId是否存在
// 		存在则根据规则生成用户id,验证这个uid是否存在,如果存在
// 			返回用户已存在
// 		如果不存在
// 			创建用户,同时给这个用户分配对应这个组的默认权限
// 	*/

// 	return
// }

// func (h *Handler) Login(ctx context.Context, req *admin.UserParams) (resp *admin.Response) {

// 	/**
// 	登录流程,先验证GroupId是否合法存在
// 		如果存在,验证用户名密码是否正确
// 			如果正确，登录成功并初始化这个用户的权限到redis
// 		如果错误
// 			提示用户名密码错误
// 	*/

// 	return
// }
