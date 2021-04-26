package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type gammaParameters struct {
	Gamma  float32
}

func ApplyGammaOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters gammaParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.Gamma(parameters.Gamma),
	)
}
