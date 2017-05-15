package yoke

import (
	"../DAO/database"
	"log"
	"net/http"
	"time"
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
		time.Sleep(time.Duration(30) * time.Second)
		current := database.GetCurrentTemp()
		set := database.GetSetTemp()
		log.Print("Current: ")
		log.Println(current)
		log.Print("set: ")
		log.Println(set)
		if current > set {
			log.Println("Lets turn on AC")
		} else if set > current {
			log.Println("Lets turn on Heater")
		} else {
			log.Println("looks like we good")
		}
	}
}
