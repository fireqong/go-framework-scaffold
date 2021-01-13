package main

import (
	"main/initialization"
	"main/kernel"
)

func main() {
	initialization.Config()
	kernel.Redis = initialization.Redis()
	kernel.App = initialization.App()
	//kernel.DB = initialization.DB()

	defer func() {
		//mysqlDB, _ := kernel.DB.DB()
		//mysqlDB.Close()

		kernel.Redis.Close()
	}()

	if err := kernel.App.Run(":12345"); err != nil {
		panic(err.Error())
	}
}
