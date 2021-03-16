package poc

import (
	"github.com/disintegration/gift"
	"strconv"
)

const RESIZE_WIDTH string = "resize-w"
const RESIZE_HEIGHT string = "resize-h"

// TODO: More operations, written only for resize
// TODO: This is just quick and dirty ._.
func applyQueryOperations(pipeline *gift.GIFT, queryArguments map[string][]string) {
	// TODO: Refactor!, when more operations are possible this gets messy and wont function
	extractedResizeWidth, resizeWidthPresent := queryArguments[RESIZE_WIDTH]
	extractedResizeHeight, resizeHeightPresent := queryArguments[RESIZE_HEIGHT]

	resizeWidth, resizeHeight := 0, 0

	if resizeWidthPresent {
		width, err := strconv.Atoi(extractedResizeWidth[0])
		if nil == err {
			resizeWidth = width
		}
	}

	if resizeHeightPresent {
		height, err := strconv.Atoi(extractedResizeHeight[0])
		if nil == err {
			resizeHeight = height
		}
	}

	if 0 != resizeWidth || 0 != resizeHeight {
		pipeline.Add(gift.Resize(resizeWidth, resizeHeight, gift.LanczosResampling))
	}
}
