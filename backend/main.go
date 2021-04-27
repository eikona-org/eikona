package main

import (
	"github.com/imgProcessing/backend/v2/data"
	"github.com/imgProcessing/backend/v2/web"
)
// @title Image Procsessing API
// @version 1.0
// @description Convert images on the fly
// @termsOfService https://github.com/imgProcessing/backend/blob/main/LICENSE

// @contact.name API Support
// @contact.url https://github.com/imgProcessing/backend/issues

// @securityDefinitions.apikey jwtAuth
// @in header
// @name Authorization

// @license.name Apache 2.0
// @license.url https://github.com/imgProcessing/backend/blob/main/LICENSE

// @BasePath /api
func main() {
	err := data.Init()
	if err != nil {
		panic(err)
	}

	web.Serve()
}
