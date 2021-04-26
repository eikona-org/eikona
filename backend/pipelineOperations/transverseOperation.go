package pipelineOperations

import (
	"github.com/disintegration/gift"
)

func ApplyTransverseOperation(pipeline *gift.GIFT) {
	pipeline.Add(
		gift.Transverse(),
	)
}
