package main

import (
	"./DAO/database"
	"./SensorReader"
	"log"
)

func main() {
	log.Println("Hello,World")
	database.InitDB()
	go SensorReader.StartSensorReader()
	log.Println("Hello,World")
	log.Println("Hello,World")
	log.Println("Hello,World")
	log.Println("Hello,World")
	log.Println("Hello,World")
	select {}
}
