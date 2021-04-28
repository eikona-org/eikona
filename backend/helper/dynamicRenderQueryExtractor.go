package helper

import (
	"fmt"
	datamodels "github.com/eikona-org/eikona/v2/data/datamodels"
	"strconv"
)

const ARGUMENT_RESIZE_WIDTH string = "resize-w"
const ARGUMENT_RESIZE_HEIGHT string = "resize-h"

type operationsConfiguration struct {
	resizeWidth  int
	resizeHeight int
}

func ExtractProcessingSteps(arguments map[string][]string) []datamodels.ProcessingStep {
	operationsConfig := operationsConfiguration{}

	for argumentName := range arguments {
		if argumentName == ARGUMENT_RESIZE_WIDTH || argumentName == ARGUMENT_RESIZE_HEIGHT {
			updateOperationsConfiguration(
				argumentName,
				&operationsConfig,
				arguments[argumentName][len(arguments[argumentName])-1],
			)
		}
	}

	return applyOperationsToPipeline(operationsConfig)
}

func applyOperationsToPipeline(operationsConfig operationsConfiguration) []datamodels.ProcessingStep {
	var processSteps []datamodels.ProcessingStep

	// Resize
	if 0 != operationsConfig.resizeWidth || 0 != operationsConfig.resizeHeight {
		processSteps = append(processSteps, datamodels.ProcessingStep{
			ProcessingStepType: datamodels.Resize,
			ParameterJson:      fmt.Sprintf(`{"width":"%d","height":"%d"}`, operationsConfig.resizeWidth, operationsConfig.resizeHeight),
		})
	}

	return processSteps
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
