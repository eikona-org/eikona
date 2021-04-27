package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type contrastSigmoidParameters struct {
	Midpoint float32
	Factor   float32
}

func ApplyContrastSigmoidOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters contrastSigmoidParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	if parameters.Midpoint > 1 {
		parameters.Midpoint = 1
	} else if parameters.Midpoint < 0 {
		parameters.Midpoint = 0
	}

	if parameters.Factor > 100 {
		parameters.Factor = 100
	} else if parameters.Factor < -100 {
		parameters.Factor = -100
	}

	pipeline.Add(
		gift.Sigmoid(parameters.Midpoint, parameters.Factor),
	)
}
