package SensorReader

import "time"
import "fmt"
import "math/rand"
import "../DAO/database"

//adding sudo data for testing

func StartSensorReader() {
	fmt.Println("Sensor reader starting")
	for {
		readSensor()
		time.Sleep(time.Duration(3) * time.Second)
	}
}
func readSensor() {
	data := rand.Intn(100)
	writeData(data)
}

func writeData(data int) {
	fmt.Println("Writing sensor reading")
	fmt.Println(data)
	database.InsertSensorData(data)
}
