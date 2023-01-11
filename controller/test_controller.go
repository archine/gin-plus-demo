package controller

import (
	"gin-plus-demo/model"
	"github.com/archine/gin-plus/v2/mvc"
	"github.com/archine/gin-plus/v2/resp"
	"github.com/gin-gonic/gin"
)

type TestController struct {
	mvc.Controller
	UserMapper *model.UserMapper `@autowired:""`
}

func init() {
	mvc.Register(&TestController{})
}

// User
// @GET(path="/user", globalFunc=false) 获取用户
func (t *TestController) User(ctx *gin.Context) {
	resp.Json(ctx, t.UserMapper.GetUser())
}

// AddUser
// @POST(path="/add_user", globalFunc=true) 添加用户
func (t *TestController) AddUser(ctx *gin.Context) {
	var arg model.User
	if resp.ParamValid(ctx, ctx.ShouldBindJSON(&arg), &arg) {
		return
	}
	t.UserMapper.AddUser(&arg)
	resp.Ok(ctx)
}
