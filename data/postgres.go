package data

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/imgProcessing/backend/v2/data/models"
)

func GetContext() *pg.DB {
	context := pg.Connect(&pg.Options{ //Requires existing DB and User TODO: Create DB on connection if it doesn't exist
		User: "imgProcessing", //TODO: Make postgres config configurable with config file or comparable
		Password: "imgProcessing",
		Addr: "127.0.0.1:5432",
		Database: "imgProcessing",
	})
	defer context.Close()

	err := createSchema(context)
	if err!= nil {
		panic(err)
	}

	return context
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*data.Image)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}