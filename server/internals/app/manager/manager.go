package manager

import (
	common_z "iads/server/pkg/common"
	"net/http"
	"time"

	webconfig "iads/server/internals/app/manager/config"
	"iads/server/internals/app/manager/controllers/common"
	"iads/server/internals/app/manager/middleware"
	"iads/server/internals/app/manager/routers"
	"iads/server/internals/pkg/config"
	"iads/server/internals/pkg/models"
	"iads/server/pkg/convert"
	"iads/server/pkg/logger"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// 运行
func Run(configPath string) {
	if configPath == "" {
		configPath = "./cfg.yaml"
	}
	// 加载配置
	cfg := webconfig.LoadDefaultConfig()

	logger.InitLog("debug", "./log/logb.log")
	initDB(cfg)
	_ = common.InitCsbinEnforcer()
	initWeb(cfg)
	logger.Debug(cfg.Web.Domain + "Api Server已启动...")
}

func initDB(config *config.Config) {
	models.InitDB(config)
	models.Migration()
}

func initWeb(config *config.Config) {
	gin.SetMode(gin.DebugMode) //调试模式
	app := gin.New()
	app.NoRoute(middleware.NoRouteHandler())
	// 崩溃恢复
	app.Use(middleware.RecoveryMiddleware())
	//app.LoadHTMLGlob(config.Web.StaticPath + "dist/*.html")
	//app.Static("/static", config.Web.StaticPath+"dist/static")
	//app.Static("/resource", config.Web.StaticPath+"resource")
	//app.StaticFile("/favicon.ico", config.Web.StaticPath+"dist/favicon.ico")
	// 注册路由
	routers.RegisterRouter(app)
	go initHTTPServer(config, app)
}

// InitHTTPServer 初始化http服务
func initHTTPServer(config *config.Config, handler http.Handler) {
	srv := &http.Server{
		Addr:         ":" + convert.ToString(config.Web.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		//IdleTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	common_z.CheckErr(err)
}
