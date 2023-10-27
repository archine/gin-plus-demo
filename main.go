package main

import (
	_ "gin-plus-demo/controller"
	"github.com/archine/gin-plus/v3/application"
)

//go:generate gp-ast
func main() {
	application.Default().Run()
}
