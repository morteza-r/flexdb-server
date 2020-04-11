package infrastructure

import (
	"../application"
	"github.com/golobby/container"
	"github.com/morteza-r/flexdb"
)

func SetUp() {

	container.Singleton(func() application.DbService {
		var db *flexdb.Database
		db = flexdb.NewDb()

		return application.DbService{
			Db: db,
		}
	})

}
