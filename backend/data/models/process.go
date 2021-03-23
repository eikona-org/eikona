package data

import uuid "github.com/satori/go.uuid"

type Process struct {
	ProcessId uuid.UUID `pg:"type:uuid,default:gen_random_uuid(),pk"`
	ProcessingSteps []ProcessingStep `pg:"rel:has-many,join_fk:process"`
}