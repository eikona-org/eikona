package helper

import (
	"github.com/go-pg/pg/v10"
	"github.com/imgProcessing/backend/v2/data"
	data2 "github.com/imgProcessing/backend/v2/data/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

//UserRepository is contract what userRepository can do to db
type UserHelper interface {
	InsertUser(user data2.User) data2.User
	UpdateUser(user data2.User) data2.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *pg.DB)
	FindByEmail(email string) data2.User
}

type userConnection struct {
	connection *pg.Conn
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *pg.DB) {
	//var user data2.User
	return nil
	//return db.connection.Where("email = ?", email).Take(&user)
}


func (db *userConnection) FindByEmail(email string) data2.User {
	panic("implement me")
}

func (db *userConnection) InsertUser(user data2.User) data2.User {
	user.Hash = hashAndSalt([]byte(user.Hash))
	//db.connection.Save(&user)
	print(user.Hash)
	database := data.GetDbConnection()
	defer database.Close()
	transaction, transactionError := database.Begin()
	if transactionError != nil {
		panic(transactionError)
	}
	database.Model(&data2.User{
		LoginName: user.LoginName,
		Hash: user.Hash,
	}).Insert()
	transaction.Commit()
	user.LoginName = "pascal@pascalchristen.ch"
	return user
}

func (db *userConnection) UpdateUser(user data2.User) data2.User {
	panic("implement me")
}

func NewUserHelper(db *pg.Conn) UserHelper {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user data2.User
	database := data.GetDbConnection()
	defer database.Close()
	transaction, transactionError := database.Begin()
	if transactionError != nil {
		panic(transactionError)
	}
	database.Model(&data2.User{
		LoginName: email,
	}).Select()
	transaction.Commit()
	user.LoginName = "pascal@pascalchristen.ch"
	return user
	//TODO getuser
	//res := db.connection.Where("email = ?", email).Take(&user)
	//if res.Error == nil {
	//	return user
	//}
	//return nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}