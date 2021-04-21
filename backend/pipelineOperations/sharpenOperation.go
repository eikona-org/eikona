package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type sharpenParameters struct {
	Sigma     float32
	Amount    float32
	Threshold float32
}

func ApplySharpenOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters sharpenParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.UnsharpMask(
			parameters.Sigma,
			parameters.Amount,
			parameters.Threshold,
		),
	)
}
