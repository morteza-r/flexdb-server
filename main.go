package main

import (
	"./app/api/rest"
	"./app/infrastructure"
	"fmt"
)

func main() {
	infrastructure.SetUp()

	err := rest.SetUp()
	if err != nil {
		fmt.Println(err)
	}
}
