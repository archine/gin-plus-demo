package controller

import (
	"gin-plus-demo/model"
	"github.com/archine/gin-plus/v3/mvc"
	"github.com/archine/gin-plus/v3/resp"
	"github.com/gin-gonic/gin"
)

// UserController
// @BasePath("/user")
type UserController struct {
	mvc.Controller
	UserMapper *model.UserMapper
}

// User
// @GET(path="/") 获取用户
func (u *UserController) User(ctx *gin.Context) {
	resp.Json(ctx, u.UserMapper.GetUser())
}

// AddUser
// @POST(path="/") 添加用户
func (u *UserController) AddUser(ctx *gin.Context) {
	var arg model.User
	if !resp.ParamValidation(ctx, &arg) {
		return
	}
	u.UserMapper.AddUser(&arg)
	resp.Ok(ctx)
}
