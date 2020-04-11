package infrastructure

import (
	"github.com/golobby/container"
	"github.com/morteza-r/flexdb"
	"github.com/morteza-r/flexdb-server/app/application"
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
