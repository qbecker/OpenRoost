package yoke

import (
	"../DAO/database"
	"log"
	"net/http"
	"time"
	"../utils"
	"../logger"
)

func StartYoke() {
	log.Println("starting Yoke")
	go monitorTemp()
	yokeContr := http.NewServeMux()
	yokeContr.HandleFunc("/settemp/", setTemp)
	yokeContr.HandleFunc("/sethometemp/", setHomeTemp)
	yokeContr.HandleFunc("/setawaytemp/", setAwayTemp)
	yokeContr.HandleFunc("/setzipcode/", setZipCode)
	yokeContr.HandleFunc("/heatoff/", heatOff)
	yokeContr.HandleFunc("/fanon/", fanOn)
	yokeContr.HandleFunc("/fanoff/", fanOff)
	yokeContr.HandleFunc("/getstatus/", getStatus)
	yokeContr.HandleFunc("/ishome/", isHome)
	yokeContr.HandleFunc("/isaway/", isAway)
	log.Fatal(http.ListenAndServe(":8082", yokeContr))
}

func monitorTemp() {
	for {
		//get current temp
		//compare to desired temp
		//turn on or off HVAC
		time.Sleep(time.Duration(3) * time.Second)
		current := database.GetCurrentTemp()
		set := database.GetSetTemp()
		log.Print("Current: ")
		log.Println(current)
		log.Print("set: ")
		log.Println(set)
		if current > set {
			log.Println("Lets turn on AC")
			utils.SendHttpRequest("POST", "Http://localhost:8081/coolon")
			logger.WriteLog("Turning on AC")
		} else if set > current {
			log.Println("Lets turn on Heater")
			utils.SendHttpRequest("POST", "Http://localhost:8081/heaton")
			logger.WriteLog("Turning on Heater")
		} else {
			log.Println("looks like we good")
		}
	}
}
