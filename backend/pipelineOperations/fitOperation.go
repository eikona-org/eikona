package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type fitParameters struct {
	Width  int
	Height int
}

func ApplyFitOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters fitParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.ResizeToFit(
			parameters.Width,
			parameters.Height,
			gift.LanczosResampling,
		),
	)
}
