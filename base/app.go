package base

import (
	"fmt"
	"gin-plus-demo/config"
	"github.com/archine/gin-plus/v2/exception"
	"github.com/archine/gin-plus/v2/plugin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type App struct {
	Engine *gin.Engine
}

// NewApp 初始化应用程序
func NewApp(configPath string) (*App, error) {
	config.InitConfig(configPath)
	initPlugin()
	engine := initGin()
	return &App{
		Engine: engine,
	}, nil
}

// Run 启动应用程序
func (a *App) Run() {
	svc := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.Port),
		Handler: a.Engine,
	}
	log.Infof("gin-plus-demo start success on Ports:[%d]", config.Conf.Port)
	if err := svc.ListenAndServe(); err != nil {
		log.Fatalf("gin-plus-demo start error, %s", err.Error())
	}
}

// 初始化全局变量
func initPlugin() {
	plugin.InitLog(config.Conf.LogLevel)
	log.Debugf("init plugin successful...")
}

// 初始化gin
func initGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(plugin.LogMiddleware())               // 日志插件
	engine.Use(exception.GlobalExceptionInterceptor) // 全局异常拦截
	engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))
	engine.MaxMultipartMemory = config.Conf.MaxFileSize
	engine.RemoveExtraSlash = true
	log.Debugf("init gin engine successful...")
	return engine
}
