package main

import (
	_ "github.com/DavidRomanovizc/Qoinify/docs"
	"github.com/DavidRomanovizc/Qoinify/internal/api"
)

// @title Qoinify API
// @version 1.0
// @description API Server for Qoinify application
func main() {
	api.Run()
}
