package helper

import (
	"github.com/go-pg/pg/v10"
	"github.com/imgProcessing/backend/v2/data"
	datamodels "github.com/imgProcessing/backend/v2/data/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserHelper interface {
	InsertUser(user datamodels.User) datamodels.User
	UpdateUser(user datamodels.User) datamodels.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *pg.DB)
	FindByEmail(email string) datamodels.User
}

type userConnection struct {
	connection *pg.DB
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *pg.DB) {
	//var user data2.User
	return nil
}


func (db *userConnection) FindByEmail(email string) datamodels.User {
	panic("implement me")
}

func (db *userConnection) InsertUser(user datamodels.User) datamodels.User {
	user.PasswordHashSalt = hashAndSalt([]byte(user.PasswordHashSalt))
	//db.connection.Save(&user)
	print(user.PasswordHashSalt)
	database := data.GetDbConnection()
	defer database.Close()
	transaction, transactionError := database.Begin()
	if transactionError != nil {
		panic(transactionError)
	}
	database.Model(&datamodels.User{
		Email: "pascal",
		PasswordHashSalt: user.PasswordHashSalt,
	}).Insert()
	transaction.Commit()
	user.Email = "pascal@pascalchristen.ch"
	return user
}

func (db *userConnection) UpdateUser(user datamodels.User) datamodels.User {
	panic("implement me")
}

func NewUserHelper(db *pg.DB) UserHelper {
	return &userConnection{
		connection: db,
	}
}


func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user datamodels.User
	database := data.GetDbConnection()
	defer database.Close()
	transaction, transactionError := database.Begin()
	if transactionError != nil {
		panic(transactionError)
	}
	database.Model(&datamodels.User{
		Email: email,
	}).Select()
	transaction.Commit()
	user.Email = "pascal@pascalchristen.ch"
	return user
	//TODO getuser
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}