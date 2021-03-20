package main

import (
	"github.com/imgProcessing/backend/v2/data"
	"github.com/imgProcessing/backend/v2/web"
)

func main() {
	err := data.Init()
	if err != nil {
		panic(err)
	}

	web.Serve()
}
