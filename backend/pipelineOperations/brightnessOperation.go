package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type brightnessParameters struct {
	Percentage float32 `json:",string"`
}

func ApplyBrightnessOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters brightnessParameters
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
		gift.Brightness(parameters.Percentage),
	)
}
