package hvaccontroller

import (
	"log"
	"net/http"
)

func StartHVACController() {
	log.Println("starting HVAC Controller")
	http.HandleFunc("/coolon", coolOn)
	http.HandleFunc("/heaton", heatOn)
	http.HandleFunc("/cooloff", coolOff)
	http.HandleFunc("/heatoff", heatOff)
	http.HandleFunc("/fanon", fanOn)
	http.HandleFunc("/fanoff", fanOff)
	http.HandleFunc("/getstatus", getStatus)
	log.Fatal(http.ListenAndServe(":9093", nil))

}
