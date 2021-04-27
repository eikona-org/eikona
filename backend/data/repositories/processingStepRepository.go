package repositories

import (
	"github.com/google/uuid"
	"github.com/eikona-org/eikona/v2/data"
	datamodels "github.com/eikona-org/eikona/v2/data/datamodels"
)

type ProcessingStepRepository struct{}

func (r ProcessingStepRepository) AddToProcess(processId uuid.UUID, stepType datamodels.ProcessingStepType, parameterJson string, executionPosition int64) error {
	dbConnection :=data.GetDbConnection()
	defer dbConnection.Close()

	transaction, transactionError := dbConnection.Begin()
	if transactionError != nil {
		return transactionError
	}

	//Shift ExecutionPositions back if necessary
	existingSteps := findProcessingStepsByProcessId(processId)
	positionBuffer := executionPosition
	for i := 0; i < len(*existingSteps); i++ {
		if (*existingSteps)[i].ExecutionPosition < executionPosition {
			continue
		}

		currentStep := (*existingSteps)[i]
		if currentStep.ExecutionPosition == positionBuffer {
			currentStep.ExecutionPosition += 1
			positionBuffer = currentStep.ExecutionPosition

			_, updateError := dbConnection.Model(&currentStep).WherePK().Update()
			if updateError != nil {
				transaction.Rollback()
				return updateError
			}
		}
	}

	_, creationError := dbConnection.Model(&datamodels.ProcessingStep{
		ProcessingStepType: stepType,
		ParameterJson: parameterJson,
		ExecutionPosition: executionPosition,
		ProcessId: processId,
	}).Insert()
	if creationError != nil {
		transaction.Rollback()
		return creationError
	}

	transaction.Commit()
	return nil
}

func (r ProcessingStepRepository) FindByProcessId(id uuid.UUID) *[]datamodels.ProcessingStep {
	return findProcessingStepsByProcessId(id)
}

func findProcessingStepsByProcessId(processId uuid.UUID) *[]datamodels.ProcessingStep {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	processingSteps := new([]datamodels.ProcessingStep)
	err := dbConnection.Model(&datamodels.ProcessingStep{}).
		Where("process_id = ?", processId).
		Order("execution_position ASC").
		Select(processingSteps)
	if err != nil {
		return nil
	}

	return processingSteps
}
