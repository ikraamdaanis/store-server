package main

import (
	"github.com/ikraamdaanis/store-server/internal/api"
	"github.com/ikraamdaanis/store-server/internal/database"
	"github.com/ikraamdaanis/store-server/internal/initializers"
)

func init() {
	initializers.LoadVariables()
	database.ConnectDatabase()
}

func main() {
	api.RunServer()
}
