package repositories

import (
	"github.com/eikona-org/eikona/v2/data"
	datamodels "github.com/eikona-org/eikona/v2/data/datamodels"
	"github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
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

func (r ProcessRepository) Create(name string, orgId uuid.UUID) *datamodels.Process {
	return createProcess(name, orgId)
}

func createProcess(name string, orgId uuid.UUID) *datamodels.Process {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	transaction, transactionError := dbConnection.Begin()
	if transactionError != nil {
		panic(transactionError)
	}
	process := &datamodels.Process{
		Name:    name,
		OwnerId: orgId,
	}
	_, creationError := dbConnection.Model(process).Insert()
	if creationError != nil {
		transaction.Rollback()
		panic(creationError)
	}
	transaction.Commit()

	return process
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
		Relation("ProcessingSteps", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("processing_step.execution_position ASC"), nil
		}).
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
		Relation("ProcessingSteps", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("processing_step.execution_position ASC"), nil
		}).
		WherePK().
		Select()

	if err != nil {
		return nil
	}

	return process
}
