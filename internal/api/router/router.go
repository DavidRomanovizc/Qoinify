package router

import (
	"fmt"
	_ "github.com/DavidRomanovizc/Qoinify/docs"
	"github.com/DavidRomanovizc/Qoinify/internal/api/controllers"
	"github.com/DavidRomanovizc/Qoinify/internal/api/middlewares"
	"github.com/DavidRomanovizc/Qoinify/pkg/orderbook"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"os"
)

func Setup() *gin.Engine {
	app := gin.New()
	ex := orderbook.NewExchange()
	// Logs
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())
	app.NoRoute(middlewares.NoMethodHandler())

	// Routers
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.GET("api/ping", controllers.GetPlaceOrder)
	app.GET("api/book/:market", controllers.HandleGetBook(ex))
	app.POST("api/order", controllers.HandlePlaceOrder(ex))
	app.DELETE("api/order/:id", controllers.HandleCancelOrder(ex))

	return app
}
