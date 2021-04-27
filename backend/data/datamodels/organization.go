package data

import (
	"github.com/google/uuid"
)

type Organization struct {
	OrganizationId	uuid.UUID	`pg:"type:uuid,default:gen_random_uuid(),pk"`
	Name			string		`pg:",notnull"`
	MinioBucketName string		`pg:",notnull"`
	Users			[]User		`pg:"rel:has-many,join_fk:organization"`
	Images			[]Image		`pg:"rel:has-many,join_fk:owner"`
	Processes		[]Process	`pg:"rel:has-many,join_fk:owner"`
}
