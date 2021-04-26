package pipelineOperations

import (
	"github.com/disintegration/gift"
)

func ApplyFlipVerticalOperation(pipeline *gift.GIFT) {
	pipeline.Add(
		gift.FlipVertical(),
	)
}
