package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type contrastParameters struct {
	Percentage float32
}

func ApplyContrastOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters contrastParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	if parameters.Percentage > 100 {
		parameters.Percentage = 100
	} else if parameters.Percentage < -100 {
		parameters.Percentage = -100
	}

	pipeline.Add(
		gift.Contrast(parameters.Percentage),
	)
}
