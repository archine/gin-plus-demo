package main

import (
	_ "gin-plus-demo/controller"
	"github.com/archine/gin-plus/v2/application"
)

//go:generate mvc
func main() {
	application.Default().Run()
}
