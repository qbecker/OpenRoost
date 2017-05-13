package hvaccontroller

import (
	"log"
	"net/http"
)

func StartHVACController() {
	HVACContr := http.NewServeMux()
	log.Println("starting HVAC Controller")
	HVACContr.HandleFunc("/coolon", coolOn)
	HVACContr.HandleFunc("/heaton", heatOn)
	HVACContr.HandleFunc("/cooloff", coolOff)
	HVACContr.HandleFunc("/heatoff", heatOff)
	HVACContr.HandleFunc("/fanon", fanOn)
	HVACContr.HandleFunc("/fanoff", fanOff)
	HVACContr.HandleFunc("/getstatus", getStatus)
	log.Fatal(http.ListenAndServe(":8081", HVACContr))
}
