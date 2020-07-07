package main

import (
	"artikel/master"
	"artikel/master/config"
)

func main() {
	db := config.Connection()
	router := config.CreateRouter()
	master.InitAll(router, db)
	config.RunServer(router)
}
