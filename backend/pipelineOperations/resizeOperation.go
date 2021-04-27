package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type resizeParameters struct {
	Width  int
	Height int
}

func ApplyResizeOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters resizeParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	if parameters.Width > 5000 {
		parameters.Width = 5000
	} else if parameters.Width < 0 {
		parameters.Width = 0
	}

	if parameters.Height > 5000 {
		parameters.Height = 5000
	} else if parameters.Height < 0 {
		parameters.Height = 0
	}

	pipeline.Add(
		gift.Resize(
			parameters.Width,
			parameters.Height,
			gift.LanczosResampling,
		),
	)
}
