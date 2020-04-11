package application

import (
	"github.com/morteza-r/flexdb"
	"time"
)

type DbService struct {
	Db *flexdb.Database
}

func (us *DbService) Query(query *flexdb.Query) (result interface{}, err error) {
	start := time.Now()
	result, err = us.Db.Run(query)
	query.Took = time.Since(start)
	if err != nil {
		return
	}

	return
}
