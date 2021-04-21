package pipelineOperations

import (
	"github.com/disintegration/gift"
)

func ApplyGrayscaleOperation(pipeline *gift.GIFT) {
	pipeline.Add(
		gift.Grayscale(),
	)
}
