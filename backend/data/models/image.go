package data

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Image struct {
	ImageId uuid.UUID `pg:"type:uuid,default:gen_random_uuid(),pk"`
	Name string
	Uploaded time.Time `pg:"default:now()"`
	MinioObjectName string
	Owner Organization `pg:"rel:has-one,fk:owner_id"`
	OwnerId uuid.UUID
}