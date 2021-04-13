package models

import (
	"github.com/google/uuid"
	"time"
)

type Image struct {
	ImageId  uuid.UUID
	Name     string
	Uploaded time.Time
}
