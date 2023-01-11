package main

import (
	"gin-plus-demo/base"
	_ "gin-plus-demo/controller"
	"github.com/archine/gin-plus/v2/ast"
	"github.com/archine/gin-plus/v2/mvc"
	log "github.com/sirupsen/logrus"
	"os"
)

//go:generate go run main.go ast
func main() {
	if len(os.Args) > 1 && os.Args[1] == "ast" {
		ast.Parse("controller")
		return
	}
	app, err := base.NewApp("app.yml")
	if err != nil {
		log.Fatalf("initializer application failed,%s", err.Error())
	}
	mvc.Apply(app.Engine, true, base.Ast)
	app.Run()
}
