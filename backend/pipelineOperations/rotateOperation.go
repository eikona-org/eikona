package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
	"image/color"
)

type rotateParameters struct {
	Angle float32
}

func ApplyRotateOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters rotateParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	pipeline.Add(
		gift.Rotate(
			parameters.Angle,
			color.White,
			gift.LinearInterpolation,
		),
	)
}
