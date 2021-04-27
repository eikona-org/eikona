package main

import (
	"github.com/eikona-org/eikona/v2/data"
	"github.com/eikona-org/eikona/v2/web"
)
// @title Image Procsessing API
// @version 1.0
// @description Convert images on the fly
// @termsOfService https://github.com/eikona-org/eikona/blob/main/LICENSE

// @contact.name API Support
// @contact.url https://github.com/eikona-org/eikona/issues

// @securityDefinitions.apikey jwtAuth
// @in header
// @name Authorization

// @license.name Apache 2.0
// @license.url https://github.com/eikona-org/eikona/blob/main/LICENSE

// @BasePath /api
func main() {
	err := data.Init()
	if err != nil {
		panic(err)
	}

	web.Serve()
}
