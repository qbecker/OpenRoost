package main

import (
	"./DAO/database"
	"./HVACController"
	"./SensorReader"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Hello,World!")
	database.InitDB()
	go SensorReader.StartSensorReader()
	go hvaccontroller.StartHVACController()
	select {}
}
