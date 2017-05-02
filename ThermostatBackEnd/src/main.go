package main

import (
        "log"
        "./SensorReader")


func main() {
	log.Println("Hello,World")
	 go SensorReader.StartSensorReader();
	 select{}
}