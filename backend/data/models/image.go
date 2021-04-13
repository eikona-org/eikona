package data

import (
	"github.com/google/uuid"
	"time"
)

type Image struct {
	ImageId         uuid.UUID    `pg:"type:uuid,default:gen_random_uuid(),pk"`
	Name            string       `pg:",notnull"`
	Uploaded        time.Time    `pg:"default:now(),notnull"`
	MinioObjectName string       `pg:",notnull"`
	Owner           Organization `pg:"rel:has-one,fk:owner_id"`
	OwnerId         uuid.UUID    `pg:"type:uuid"`
}