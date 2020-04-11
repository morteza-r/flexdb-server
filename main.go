package main

import (
	"fmt"
	"github.com/morteza-r/flexdb-server/app/api/rest"
	"github.com/morteza-r/flexdb-server/app/infrastructure"
)

func main() {
	infrastructure.SetUp()

	err := rest.SetUp()
	if err != nil {
		fmt.Println(err)
	}
}
