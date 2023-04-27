package api

import (
	"fmt"
	"github.com/DavidRomanovizc/Qoinify/internal/api/router"
)

func setConfiguration() {
	fmt.Println("Set configuration...")
}

func Run() {
	setConfiguration()
	web := router.Setup()
	port := "8080"
	fmt.Println("Go API REST Running on port" + port)
	fmt.Println("==================>")
	_ = web.Run(":" + port)
}
