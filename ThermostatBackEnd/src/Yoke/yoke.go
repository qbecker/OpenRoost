package yoke

import (
	"log"
	"net/http"
)

func StartYoke() {
	log.Println("starting Yoke")
	yokeContr := http.NewServeMux()
	yokeContr.HandleFunc("/settemp/", setTemp)
	yokeContr.HandleFunc("/sethometemp", coolOn)
	yokeContr.HandleFunc("/setawaytemp", heatOn)
	yokeContr.HandleFunc("/setzipcode", coolOff)
	yokeContr.HandleFunc("/heatoff", heatOff)
	yokeContr.HandleFunc("/fanon", fanOn)
	yokeContr.HandleFunc("/fanoff", fanOff)
	yokeContr.HandleFunc("/getstatus", getStatus)
	log.Fatal(http.ListenAndServe(":8082", yokeContr))
}
