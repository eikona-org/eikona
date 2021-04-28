package pipelineOperations

import (
	"encoding/json"
	"github.com/disintegration/gift"
	"image/color"
	"math"
)

type rotateParameters struct {
	Angle float32 `json:",string"`
}

func ApplyRotateOperation(pipeline *gift.GIFT, params string) {
	b := []byte(params)
	var parameters rotateParameters
	err := json.Unmarshal(b, &parameters)

	if err != nil {
		return
	}

	if parameters.Angle > 360 {
		parameters.Angle = float32(math.Mod(float64(parameters.Angle), 360))
	} else if parameters.Angle < -360 {
		parameters.Angle = float32(math.Mod(float64(parameters.Angle), -360))
	}

	pipeline.Add(
		gift.Rotate(
			parameters.Angle,
			color.White,
			gift.LinearInterpolation,
		),
	)
}
