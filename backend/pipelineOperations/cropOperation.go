package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type cropParameters struct {
	Width  int
	Height int
}

func ApplyCropOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters cropParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.CropToSize(
			parameters.Width,
			parameters.Height,
			gift.TopLeftAnchor,
		),
	)
}
