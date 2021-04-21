package pipelineOperations

import (
	"github.com/disintegration/gift"
)

func ApplyInvertOperation(pipeline *gift.GIFT) {
	pipeline.Add(
		gift.Invert(),
	)
}
