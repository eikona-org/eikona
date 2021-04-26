package pipelineOperations

import (
	"github.com/disintegration/gift"
)

func ApplyFlipHorizontalOperation(pipeline *gift.GIFT) {
	pipeline.Add(
		gift.FlipHorizontal(),
	)
}
