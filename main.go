package main

import (
	_ "gin-plus-demo/controller"
	"github.com/archine/gin-plus/v2/application"
)

//go:generate mvc2
func main() {
	application.Default().Run()
}
