package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type cropCenterParameters struct {
	Width  int
	Height int
}

func ApplyCropCenterOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters cropCenterParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.CropToSize(
			parameters.Width,
			parameters.Height,
			gift.CenterAnchor,
		),
	)
}
