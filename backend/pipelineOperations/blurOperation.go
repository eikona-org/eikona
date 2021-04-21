package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type blurParameters struct {
	Sigma  float32
}

func ApplyBlurOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters blurParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.GaussianBlur(parameters.Sigma),
	)
}
