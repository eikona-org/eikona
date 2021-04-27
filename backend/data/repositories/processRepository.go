package repositories

import (
	"github.com/google/uuid"
	"github.com/imgProcessing/backend/v2/data"
	datamodels "github.com/imgProcessing/backend/v2/data/datamodels"
)

type ProcessRepository struct{}

func (r ProcessRepository) GetAll(orgId uuid.UUID) *[]datamodels.Process {
	return getAllProcesses(orgId)
}

func (r ProcessRepository) Find(id uuid.UUID) *datamodels.Process {
	return findProcess(id)
}

func (r ProcessRepository) FindByIdAndOrganizationId(id uuid.UUID, orgId uuid.UUID) *datamodels.Process {
	return findProcessByIdAndOrganizationId(id, orgId)
}

func getAllProcesses(orgId uuid.UUID) *[]datamodels.Process {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	processes := new([]datamodels.Process)
	err := dbConnection.Model(&datamodels.Process{}).
		Column("process_id").
		Column("name").
		Where("owner_id = ?", orgId).
		Select(processes)
	if err != nil {
		return nil
	}

	return processes
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

func findProcessByIdAndOrganizationId(id uuid.UUID, ownerId uuid.UUID) *datamodels.Process {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	process := &datamodels.Process{ProcessId: id, OwnerId: ownerId}
	err := dbConnection.Model(process).
		Relation("ProcessingSteps").
		WherePK().
		Select()

	if err != nil {
		return nil
	}

	return process
}
