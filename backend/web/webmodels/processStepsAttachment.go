package webmodels

type ProcessStepsAttachment struct {
	ProcessingSteps []ProcessStepAttachment `form:"processingSteps" json:"processingSteps" binding:"required" example:"{
	"processingSteps": [
		{
			"processId" : "b2058e34-d8fa-44cd-af75-1983a0c5f172",
			"processingStepType" : 40,
			"executionPosition" : 2,
			"parameterJson" : "{}"
		},{
			"processId" : "b2058e34-d8fa-44cd-af75-1983a0c5f172",
			"processingStepType" : 60,
			"executionPosition" : 3,
			"parameterJson" : "{ 'var': 'val' }"
		}
	]
}"`
}