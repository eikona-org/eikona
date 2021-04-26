package pipelineOperations

import (
	"github.com/disintegration/gift"
)

func ApplyTransposeOperation(pipeline *gift.GIFT) {
	pipeline.Add(
		gift.Transpose(),
	)
}
