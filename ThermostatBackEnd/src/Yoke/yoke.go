package yoke

import (
	"log"
	"net/http"
	"time"
)

func StartYoke() {
	log.Println("starting Yoke")
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

func monitorTemp(){
	for{
		//get current temp
		//compare to desired temp
		//turn on or off HVAC
		
	}
}
