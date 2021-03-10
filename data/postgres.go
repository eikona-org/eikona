package data

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/imgProcessing/backend/v2/data/models"
	"os"
)

func GetContext() *pg.DB {
	opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	defer db.Close()

	err = createSchema(db)
	if err != nil {
		panic(err)
	}

	return db
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