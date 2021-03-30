package data

import (
	"github.com/google/uuid"
)

type ProcessingStep struct {
	ProcessingStepId uuid.UUID `pg:"type:uuid,default:gen_random_uuid(),pk"`
	ProcessingStepType ProcessingStepType
	ParameterJson string
	ExecutionPosition int64
	Process Process `pg:"rel:has-one,fk:process_id"`
	ProcessId uuid.UUID
}