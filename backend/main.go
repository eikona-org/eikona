package main

import (
	"github.com/imgProcessing/backend/v2/data"
	"github.com/imgProcessing/backend/v2/web"
)

func main() {
	context := data.GetContext()

	web.Serve(context)
}
