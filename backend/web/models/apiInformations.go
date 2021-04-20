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

type Process struct {
	ProcessId  uuid.UUID
	Name     string
}
