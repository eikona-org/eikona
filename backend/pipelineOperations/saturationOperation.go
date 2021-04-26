package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type saturationParameters struct {
	Percentage float32
}

func ApplySaturationOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters saturationParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.Saturation(parameters.Percentage),
	)
}
