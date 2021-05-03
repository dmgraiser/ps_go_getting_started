package main

import (
	"fmt"

	"github.com/dmgraiser/ps_go_getting_started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Doug",
		LastName:  "Graiser",
	}
	fmt.Println(u)
}
