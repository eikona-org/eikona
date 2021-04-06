package repositories

import (
	"github.com/google/uuid"
	"github.com/imgProcessing/backend/v2/data"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
)

type ImageRepository struct{}

func (r ImageRepository) Find(id uuid.UUID, orgId uuid.UUID) *datamodels.Image {
	return findImage(id, orgId)
}

func findImage(id uuid.UUID, orgId uuid.UUID) *datamodels.Image {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	var image datamodels.Image
	err := dbConnection.Model(&datamodels.Image{
		ImageId: id,
		OwnerId: orgId,
	}).Select(image)
	if err != nil {
		return nil
	}

	return &image
}
