package repositories

import (
	"github.com/google/uuid"
	"github.com/imgProcessing/backend/v2/data"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	webmodels "github.com/imgProcessing/backend/v2/web/models"
)

type ImageRepository struct{}

func (r ImageRepository) Find(id uuid.UUID, orgId uuid.UUID) *datamodels.Image {
	return findImage(id, orgId)
}

func (r ImageRepository) AllImages(orgId uuid.UUID) []webmodels.Image {
	return getAllImages(orgId)
}

func getAllImages(orgId uuid.UUID) []webmodels.Image {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	var imageTest []webmodels.Image
	err := dbConnection.Model(&datamodels.Image{}).
		Column("image_id").
		Column("name").
		Column("uploaded").
		Where("owner_id = ?", orgId).
		Select(&imageTest)
	if err != nil {
		return nil
	}
	return imageTest
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
