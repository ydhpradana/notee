package main

import (
	"notee/configs"
	"notee/routes"
)


func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8010")
}

