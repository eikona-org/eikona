package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type sharpenParameters struct {
	Sigma     float32 `json:",string"`
	Amount    float32 `json:",string"`
	Threshold float32 `json:",string"`
}

func ApplySharpenOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters sharpenParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	if parameters.Sigma > 100 {
		parameters.Sigma = 100
	} else if parameters.Sigma < 0 {
		parameters.Sigma = 0
	}

	if parameters.Amount > 10 {
		parameters.Amount = 10
	} else if parameters.Amount < 0 {
		parameters.Amount = 0
	}

	if parameters.Threshold > 1 {
		parameters.Threshold = 1
	} else if parameters.Threshold < 0 {
		parameters.Threshold = 0
	}

	pipeline.Add(
		gift.UnsharpMask(
			parameters.Sigma,
			parameters.Amount,
			parameters.Threshold,
		),
	)
}
