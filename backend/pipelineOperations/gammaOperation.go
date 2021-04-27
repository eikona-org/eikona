package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type gammaParameters struct {
	Gamma float32
}

func ApplyGammaOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters gammaParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	if parameters.Gamma > 100 {
		parameters.Gamma = 100
	} else if parameters.Gamma < 0 {
		parameters.Gamma = 0
	}

	pipeline.Add(
		gift.Gamma(parameters.Gamma),
	)
}
