package pipelineOperations

import (
	"github.com/disintegration/gift"
)

func ApplyRotate270Operation(pipeline *gift.GIFT) {
	pipeline.Add(
		gift.Rotate270(),
	)
}
