package main

import (
	"fmt"

	"github.com/Rubicosaur/golang-rest-api/models"
)

func main() {
	t := models.GenerateTicket("123", "Abc", "Abc", "Abc", "abc")
	fmt.Println(t)
}
