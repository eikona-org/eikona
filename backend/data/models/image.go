package data

import (
	"fmt"
	"time"
)

type Image struct {
	Id int64
	DateUploaded time.Time
	Hash string
	Location string
}

func (i Image) String() string {
	return fmt.Sprintf("Image<%d %s %s %s>", i.Id, i.Hash, i.Location, i.DateUploaded.String())
}
