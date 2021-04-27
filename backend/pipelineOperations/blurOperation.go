package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type blurParameters struct {
	Sigma float32
}

func ApplyBlurOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters blurParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	if parameters.Sigma > 100 {
		parameters.Sigma = 100
	} else if parameters.Sigma < 0 {
		parameters.Sigma = 0
	}

	pipeline.Add(
		gift.GaussianBlur(parameters.Sigma),
	)
}
