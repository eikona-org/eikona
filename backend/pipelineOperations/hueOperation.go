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

	pipeline.Add(
		gift.Hue(parameters.Shift),
	)
}
