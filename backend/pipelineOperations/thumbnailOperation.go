package pipelineOperations

import (
	"github.com/disintegration/gift"
)

func ApplyThumbnailOperation(pipeline *gift.GIFT) {
	pipeline.Add(
		gift.ResizeToFill(
			1280,
			720,
			gift.LanczosResampling,
			gift.CenterAnchor,
		),
	)
}
