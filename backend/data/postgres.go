package data

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/imgProcessing/backend/v2/data/datamodels"
	"os"
)

var db *pg.DB

func Init() error {
	opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db = pg.Connect(opt)

	err = createSchema()
	if err != nil {
		panic(err)
	}

	return nil
}

func GetDbConnection() *pg.Conn {
	return db.Conn()
}

func createSchema() error {
	fmt.Println("Creating DB Schema...")
	db := GetDbConnection()
	defer db.Close()

	transaction, transactionError := db.Begin()
	if transactionError != nil {
		panic(transactionError)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Something went wrong during Schema creation: %s\n", r)
			transaction.Rollback()
			panic(r)
		}
	}()

	models := []interface{}{
		(*data.Image)(nil),
		(*data.User)(nil),
		(*data.Organization)(nil),
		(*data.Process)(nil),
		(*data.ProcessingStep)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			fmt.Println("Error occurred! Rolling back...")
			transaction.Rollback()
			return err
		}
	}

	transaction.Commit()
	fmt.Println("DB Schema created!")

	err := healthCheck()
	if err != nil {
		return err
	}
	return nil
}

func healthCheck() error {
	fmt.Println("Performing DB Health Check...")
	db := GetDbConnection()

	context := db.Context()
	err := db.Ping(context)
	if err != nil {
		panic(err)
	}

	transaction, transactionError := db.Begin()
	if transactionError != nil {
		panic(transactionError)
	}
	defer db.Close()

	var users []data.User
	err = db.Model(&users).Select()
	if err != nil {
		panic(err)
	}

	var images []data.Image
	err = db.Model(&images).Select()
	if err != nil {
		panic(err)
	}

	var organizations []data.Organization
	err = db.Model(&organizations).Select()
	if err != nil {
		panic(err)
	}

	var processes []data.Process
	err = db.Model(&processes).Select()
	if err != nil {
		panic(err)
	}

	var processingSteps []data.ProcessingStep
	err = db.Model(&processingSteps).Select()
	if err != nil {
		panic(err)
	}

	transaction.Commit()

	fmt.Println("DB Health Check succeeded!")

	return nil
}
