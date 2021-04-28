package webmodels

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