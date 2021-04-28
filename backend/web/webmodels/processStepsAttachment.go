package webmodels

type ProcessStepsAttachment struct {
	ProcessingSteps []ProcessStepAttachment `form:"processingSteps" json:"processingSteps" binding:"required" example:"[{ "processId": "ecc20406-a798-11eb-bcbc-0242ac130002", "processingStepType": 10, "executionPosition": 1, "parameterJson": "{ "param": "val" }" }]"`
}