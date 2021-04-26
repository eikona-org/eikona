package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type fillParameters struct {
	Width  int
	Height int
}

func ApplyFillOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters fillParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.ResizeToFill(
			parameters.Width,
			parameters.Height,
			gift.LanczosResampling,
			gift.CenterAnchor,
		),
	)
}
