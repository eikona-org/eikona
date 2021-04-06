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

	var process datamodels.Process
	err := dbConnection.Model(&datamodels.Process{
		ProcessId: id,
	}).Select(process)
	if err != nil {
		return nil
	}

	return &process
}
