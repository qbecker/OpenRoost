package yoke

import (
	"log"
	"net/http"
)

func StartYoke() {
	log.Println("starting Yoke")
	http.HandleFunc("/settemp", setTemp)
	http.HandleFunc("/sethometemp", coolOn)
	http.HandleFunc("/setawaytemp", heatOn)
	http.HandleFunc("/setzipcode", coolOff)
	http.HandleFunc("/heatoff", heatOff)
	http.HandleFunc("/fanon", fanOn)
	http.HandleFunc("/fanoff", fanOff)
	http.HandleFunc("/getstatus", getStatus)
	log.Fatal(http.ListenAndServe(":9093", nil))
}
