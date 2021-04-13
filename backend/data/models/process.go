package data

import (
	"github.com/google/uuid"
)

type Process struct {
	ProcessId       uuid.UUID        `pg:"type:uuid,default:gen_random_uuid(),pk"`
	Name            string           `pg:",notnull"`
	ProcessingSteps []ProcessingStep `pg:"rel:has-many,join_fk:process_id"`
}
