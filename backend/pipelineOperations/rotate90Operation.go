package pipelineOperations

import (
	"github.com/disintegration/gift"
)

func ApplyRotate90Operation(pipeline *gift.GIFT) {
	pipeline.Add(
		gift.Rotate90(),
	)
}
