package data

import (
	"github.com/google/uuid"
)

type User struct {
	UserId uuid.UUID `pg:"type:uuid,default:gen_random_uuid(),pk"`
	LoginName string
	Organization Organization `pg:"rel:has-one,fk:organization_id"`
	OrganizationId uuid.UUID
}