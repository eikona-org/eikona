package webmodels

import (
	datamodels "github.com/eikona-org/eikona/v2/data/datamodels"
	"github.com/google/uuid"
)

type ProcessStepAttachment struct {
	ProcessId uuid.UUID `form:"processId" json:"processId" binding:"required" example:"ecc20406-a798-11eb-bcbc-0242ac130002"`
	ProcessingStepType datamodels.ProcessingStepType `form:"processingStepType" json:"processingStepType" binding:"required" example:"10"`
	ExecutionPosition int64 `form:"executionPosition" json:"executionPosition" binding:"required" example:"1"`
	ParameterJson string `form:"parameterJson" json:"parameterJson" binding:"required" example:"{ "param": "value" }"`
}