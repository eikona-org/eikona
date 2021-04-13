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

func (r ImageRepository) AllImages(orgId uuid.UUID) []datamodels.Image {
	return getAllImages(orgId)
}

func getAllImages(orgId uuid.UUID) []datamodels.Image {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	var image []datamodels.Image
	err := dbConnection.Model(&image).
		Where("owner_id = ?", orgId).
		Select()
	if err != nil {
		return nil
	}
	return image
}

func findImage(id uuid.UUID, orgId uuid.UUID) *datamodels.Image {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	image := &datamodels.Image{ImageId: id}
	err := dbConnection.Model(image).
		Relation("Owner").
		WherePK().
		Where("owner_id = ?", orgId).
		Select()

	if err != nil {
		return nil
	}

	return image
}
