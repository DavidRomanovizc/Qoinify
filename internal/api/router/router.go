package router

import (
	"fmt"
	"github.com/DavidRomanovizc/Qoinify/internal/api/controllers"
	"github.com/DavidRomanovizc/Qoinify/internal/api/middlewares"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Setup() *gin.Engine {
	app := gin.New()

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

	app.GET("api/ping", controllers.GetPlaceOrder)

	return app
}
