package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type hueParameters struct {
	Shift float32
}

func ApplyHueOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters hueParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	if parameters.Shift > 180 {
		parameters.Shift = 180
	} else if parameters.Shift < -180 {
		parameters.Shift = -180
	}

	pipeline.Add(
		gift.Hue(parameters.Shift),
	)
}
