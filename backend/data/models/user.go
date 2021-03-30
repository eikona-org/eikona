package data

import (
	"github.com/google/uuid"
)

type User struct {
	UserId uuid.UUID `pg:"type:uuid,default:gen_random_uuid(),pk"`
	Email string `pg:",notnull,unique"`
	PasswordHashSalt string `pg:",notnull"`
	Organization Organization `pg:"rel:has-one,fk:organization_id,notnull"`
	OrganizationId uuid.UUID `pg:"type:uuid"`
}
