package repositories

import (
	"github.com/google/uuid"
	"github.com/imgProcessing/backend/v2/data"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
)

type ProcessRepository struct{}

func (r ProcessRepository) Find(id uuid.UUID) *datamodels.Process {
	return findProcess(id)
}

func findProcess(id uuid.UUID) *datamodels.Process {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	process := &datamodels.Process{ProcessId: id}
	err := dbConnection.Model(process).
		Relation("ProcessingSteps").
		WherePK().
		Select()

	if err != nil {
		return nil
	}
	return process
}
