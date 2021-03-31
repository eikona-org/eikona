package repositories

import (
	"github.com/google/uuid"
	"github.com/imgProcessing/backend/v2/data"
	models "github.com/imgProcessing/backend/v2/data/models"
)

type OrganizationRepository struct {}

func (r OrganizationRepository) Find(id uuid.UUID) *models.Organization {
	return findOrganization(id)
}

func (r OrganizationRepository) CreateNew(name string) *models.Organization {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	transaction, transactionError := dbConnection.Begin()
	if transactionError != nil {
		panic(transactionError)
	}

	minioBucketName := uuid.NewString()
	_, creationError := dbConnection.Model(&models.Organization{
		Name:            name,
		MinioBucketName: minioBucketName,
	}).Insert()
	if creationError != nil {
		transaction.Rollback()
		panic(creationError)
	}

	organization := &models.Organization{}
	findError := dbConnection.Model(organization).Where("minio_bucket_name = ?", minioBucketName).Select()
	if findError != nil {
		transaction.Rollback()
		panic(findError)
	}

	transaction.Commit()

	return organization
}

func findOrganization(id uuid.UUID) *models.Organization {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	var organization models.Organization
	err := dbConnection.Model(&models.Organization{
		OrganizationId: id,
	}).Select(organization)
	if err != nil {
		return nil
	}

	return &organization
}