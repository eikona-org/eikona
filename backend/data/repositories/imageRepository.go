package repositories

import (
	"github.com/google/uuid"
	"github.com/imgProcessing/backend/v2/data"
	datamodels "github.com/imgProcessing/backend/v2/data/datamodels"
	"time"
)

type ImageRepository struct{}

func (r ImageRepository) GetAll(orgId uuid.UUID) *[]datamodels.Image {
	return getAllImages(orgId)
}

func (r ImageRepository) Find(id uuid.UUID) *datamodels.Image {
	return findImage(id)
}

func (r ImageRepository) FindWithOrganization(id uuid.UUID) *datamodels.Image {
	return findImageWithOrganization(id)
}

func (r ImageRepository) FindByUserId(id uuid.UUID) *[]datamodels.Image {
	return findImagesByUserId(id)
}

func (r ImageRepository) FindByOrganizationId(id uuid.UUID) *[]datamodels.Image {
	return findImagesByOrganizationId(id)
}

func (r ImageRepository) Insert(name string, minioName string, userId uuid.UUID) *datamodels.Image {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	transaction, transactionError := dbConnection.Begin()
	if transactionError != nil {
		panic(transactionError)
	}

	_, creationError := dbConnection.Model(&datamodels.Image{
		Name:            name,
		Uploaded:        time.Now(),
		MinioObjectName: minioName,
		OwnerId:         userId,
	}).Insert()
	if creationError != nil {
		transaction.Rollback()
		panic(creationError)
	}

	image := &datamodels.Image{}
	findError := dbConnection.Model(image).
		Where("minio_object_name = ?", minioName).
		Select()
	if findError != nil {
		transaction.Rollback()
		panic(findError)
	}

	transaction.Commit()

	return image
}

func getAllImages(orgId uuid.UUID) *[]datamodels.Image {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	images := new([]datamodels.Image)
	err := dbConnection.Model(&datamodels.Image{}).
		Column("image_id").
		Column("name").
		Column("uploaded").
		Where("owner_id = ?", orgId).
		Select(images)
	if err != nil {
		return nil
	}

	return images
}

func findImage(id uuid.UUID) *datamodels.Image {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	image := &datamodels.Image{ImageId: id}
	err := dbConnection.Model(image).
		Select()
	if err != nil {
		return nil
	}

	return image
}
func findImageWithOrganization(id uuid.UUID) *datamodels.Image {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	image := &datamodels.Image{ImageId: id}
	err := dbConnection.Model(image).
		Relation("Owner").
		WherePK().
		Select()

	if err != nil {
		return nil
	}

	return image
}

func findImagesByOrganizationId(id uuid.UUID) *[]datamodels.Image {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	var images = new([]datamodels.Image)
	err := dbConnection.Model(images).
		Where("owner_id = ?", id).
		Select()

	if err != nil {
		return nil
	}

	return images
}

func findImagesByUserId(id uuid.UUID) *[]datamodels.Image {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	var images = new([]datamodels.Image)
	err := dbConnection.Model(images).
		Join("inner join users u on u.user_id = ?", id).
		Where("owner_id = u.organization_id").
		Select()

	if err != nil {
		return nil
	}

	return images
}