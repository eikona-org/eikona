package pipelineOperations

import (
	"github.com/disintegration/gift"
)

func ApplyRotate180Operation(pipeline *gift.GIFT) {
	pipeline.Add(
		gift.Rotate180(),
	)
}
