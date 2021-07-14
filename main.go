package main

import (
	"github.com/crowdeco/demo-user-service/app"
	"github.com/crowdeco/demo-user-service/database"
)

func init() {
	database.Connect()
}

func main() {
	app.Run()
}
