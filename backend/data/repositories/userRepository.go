package repositories

import (
	"github.com/google/uuid"
	"github.com/imgProcessing/backend/v2/data"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func (r UserRepository) Find(id uuid.UUID) *datamodels.User {
	return findUser(id)
}

func (r UserRepository) FindByEmail(email string) *datamodels.User {
	return findUserByEmail(email)
}

func (r UserRepository) FindByOrganizationId(id uuid.UUID) *[]datamodels.User {
	return findUserByOrganizationId(id)
}

func (r UserRepository) IsDuplicateEmail(email string) bool {
	return findUserByEmail(email) == nil
}

func (r UserRepository) InsertOrUpdate(email string, password []byte, organizationId uuid.UUID) (*datamodels.User, bool) {
	oldUser := findUserByEmail(email)
	passwordHash := hashAndSalt(password)
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	transaction, transactionError := dbConnection.Begin()
	if transactionError != nil {
		panic(transactionError)
	}

	userExists := oldUser != nil
	if userExists {
		_, updateError := dbConnection.Model(&datamodels.User{
			UserId:           oldUser.UserId,
			Email:            email,
			PasswordHashSalt: passwordHash,
			OrganizationId:   organizationId,
		}).Where("email = ?", email).Update()
		if updateError != nil {
			transaction.Rollback()
			panic(updateError)
		}
	} else {
		_, insertError := dbConnection.Model(&datamodels.User{
			Email:            email,
			PasswordHashSalt: passwordHash,
			OrganizationId:   organizationId,
		}).Insert()
		if insertError != nil {
			transaction.Rollback()
			panic(insertError)
		}
	}

	transaction.Commit()

	return findUserByEmail(email), userExists
}

func (r UserRepository) VerifyCredential(email string, password string) interface{} {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	user := findUserByEmail(email)
	if user != nil {
		return user
	}
	return nil
}

func findUser(id uuid.UUID) *datamodels.User {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	user := &datamodels.User{UserId: id}
	err := dbConnection.Model(user).WherePK().Select()
	if err != nil {
		return nil
	}

	return user
}

func findUserByEmail(email string) *datamodels.User {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	user := &datamodels.User{}
	err := dbConnection.Model(user).Where("email = ?", email).Select()
	if err != nil {
		return nil
	}
	return user
}

func findUserByOrganizationId(id uuid.UUID) *[]datamodels.User {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	var users []datamodels.User
	err := dbConnection.Model(&users).Where("organization_id = ?", id).Select()
	if err != nil {
		return nil
	}

	return &users
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic("Failed to hash a password")
	}
	return string(hash)
}
