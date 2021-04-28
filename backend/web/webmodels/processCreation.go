package webmodels

import (
	datamodels "github.com/eikona-org/eikona/v2/data/datamodels"
)

type CreateProcess struct {
	Name string `form:"name" json:"name" binding:"required" example:"Test Process"`
	ProcessingSteps []ProcessStepAttachment `form:"processingSteps" json:"processingSteps" binding:"required" example:"{
	"processingSteps": [
		{
			"processingStepType" : 40,
			"executionPosition" : 2,
			"parameterJson" : "{}"
		},{
			"processingStepType" : 60,
			"executionPosition" : 3,
			"parameterJson" : "{ 'var': 'val' }"
		}
	]
}"`
}

type ProcessStepAttachment struct {
	ProcessingStepType datamodels.ProcessingStepType `form:"processingStepType" json:"processingStepType" binding:"required" example:"10"`
	ExecutionPosition int64 `form:"executionPosition" json:"executionPosition" binding:"required" example:"1"`
	ParameterJson string `form:"parameterJson" json:"parameterJson" binding:"required" example:"{ "param": "value" }"`
}
