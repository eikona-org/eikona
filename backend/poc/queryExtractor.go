package poc

import (
	"github.com/disintegration/gift"
	"sort"
	"strconv"
)

const ARGUMENT_RESIZE_WIDTH string = "resize-w"
const ARGUMENT_RESIZE_HEIGHT string = "resize-h"

type operationsConfiguration struct {
	resizeWidth  int
	resizeHeight int
}

// Use initializer function as go doesn't support const arrays
func getSupportedArguments() []string {
	return []string{
		ARGUMENT_RESIZE_WIDTH,
		ARGUMENT_RESIZE_HEIGHT,
	}
}

// TODO: More operations, written only for resize
func applyQueryOperations(imageTransformer *gift.GIFT, queryArguments map[string][]string) {
	imageTransformer.Add(createPipelineConfiguration(queryArguments)...)
}

// Query parameter can have the same name => []string , we take the last one provided
func createPipelineConfiguration(arguments map[string][]string) []gift.Filter {
	operationsConfig := operationsConfiguration{}

	for argumentName := range arguments {
		index := sort.SearchStrings(getSupportedArguments(), argumentName)

		if index < len(getSupportedArguments()) {
			updateOperationsConfiguration(
				argumentName,
				&operationsConfig,
				arguments[argumentName][len(arguments[argumentName])-1],
			)
		}
	}

	var filters []gift.Filter

	applyOperationsToPipeline(
		operationsConfig,
		&filters,
	)

	return filters
}

func applyOperationsToPipeline(operationsConfig operationsConfiguration, filters *[]gift.Filter) {
	// Resize
	if 0 != operationsConfig.resizeWidth || 0 != operationsConfig.resizeHeight {
		gift.Resize(
			operationsConfig.resizeHeight,
			operationsConfig.resizeHeight,
			gift.LanczosResampling,
		)
	}
}

func updateOperationsConfiguration(argument string, operationsConfig *operationsConfiguration, value string) {
	switch argument {
	case ARGUMENT_RESIZE_WIDTH:
		width, err := strconv.Atoi(value)
		if nil == err {
			operationsConfig.resizeWidth = width
		}
	case ARGUMENT_RESIZE_HEIGHT:
		height, err := strconv.Atoi(value)
		if nil == err {
			operationsConfig.resizeHeight = height
		}
	}
}
