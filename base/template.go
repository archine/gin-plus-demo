package base

// 自动生成,请不要编辑

import "github.com/archine/gin-plus/v2/ast"

var Ast = map[string][]*ast.MethodInfo{
	"AddUser": {
		{Method: "POST", ApiPath: "/add_user", GlobalFunc: true},
	},
	"User": {
		{Method: "GET", ApiPath: "/user", GlobalFunc: false},
	},
}
