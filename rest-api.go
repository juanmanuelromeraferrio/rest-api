package main

import (
	"github.com/juanmanuelromeraferrio/rest-api/config/database"
	"github.com/juanmanuelromeraferrio/rest-api/route"
)

func main() {
	database.Load()
	route.Run()
}
