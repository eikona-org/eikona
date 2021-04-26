package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type contrastSigmoidParameters struct {
	Midpoint float32
	Factor float32
}

func ApplyContrastSigmoidOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters contrastSigmoidParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.Sigmoid(parameters.Midpoint, parameters.Factor),
	)
}
