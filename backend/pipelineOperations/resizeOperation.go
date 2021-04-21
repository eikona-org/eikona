package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
)

type resizeParameters struct {
	Width  int
	Height int
}

func ApplyResizeOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters resizeParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.Resize(
			parameters.Height,
			parameters.Width,
			gift.LanczosResampling,
		),
	)
}
