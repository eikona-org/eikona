package data

import (
	"github.com/google/uuid"
)

type Organization struct {
	OrganizationId uuid.UUID `pg:"type:uuid,default:gen_random_uuid(),pk"`
	Name string
	MinioBucketName string
	Users []User `pg:"rel:has-many,join_fk:organization"`
	Images []Image `pg:"rel:has-many,join_fk:owner"`
}