package repositories

import (
	"github.com/google/uuid"
	"github.com/imgProcessing/backend/v2/data"
	datamodels "github.com/imgProcessing/backend/v2/data/datamodels"
)

type OrganizationRepository struct{}

func (r OrganizationRepository) Find(id uuid.UUID) *datamodels.Organization {
	return findOrganization(id)
}

func (r OrganizationRepository) CreateNew(name string) *datamodels.Organization {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	transaction, transactionError := dbConnection.Begin()
	if transactionError != nil {
		panic(transactionError)
	}

	minioBucketName := uuid.NewString()
	_, creationError := dbConnection.Model(&datamodels.Organization{
		Name:            name,
		MinioBucketName: minioBucketName,
	}).Insert()

	if creationError != nil {
		transaction.Rollback()
		panic(creationError)
	}

	organization := &datamodels.Organization{}
	findError := dbConnection.Model(organization).
		Where("minio_bucket_name = ?", minioBucketName).
		Select()

	if findError != nil {
		transaction.Rollback()
		panic(findError)
	}

	transaction.Commit()

	return organization
}

func findOrganization(id uuid.UUID) *datamodels.Organization {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	organization := &datamodels.Organization{OrganizationId: id}
	err := dbConnection.Model(organization).WherePK().
		Select(organization)

	if err != nil {
		return nil
	}

	return organization
}
