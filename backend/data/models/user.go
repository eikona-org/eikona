package data

import uuid "github.com/satori/go.uuid"

type User struct {
	UserId uuid.UUID `pg:"type:uuid,default:gen_random_uuid(),pk"`
	LoginName string
	Organization Organization `pg:"rel:has-one,fk:organization_id"`
	OrganizationId uuid.UUID
}