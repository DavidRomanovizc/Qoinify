package api

import (
	"fmt"
	"github.com/DavidRomanovizc/Qoinify/internal/api/router"
	"go.uber.org/zap"
	"log"
)

func setConfiguration() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()
	sugar.Info("Set configuration...")
}

func Run() {
	setConfiguration()
	web := router.Setup()
	port := "8080"
	fmt.Println("Go API REST Running on port " + port)
	fmt.Println("==================>")
	_ = web.Run(":" + port)
}
