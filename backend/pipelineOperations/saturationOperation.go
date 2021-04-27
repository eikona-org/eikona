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

	if parameters.Percentage > 500 {
		parameters.Percentage = 500
	} else if parameters.Percentage < -100 {
		parameters.Percentage = -100
	}

	pipeline.Add(
		gift.Saturation(parameters.Percentage),
	)
}
