package hvaccontroller

import (
	"io"
	"log"
	"net/http"
)

func coolOn(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn on cooling")
	io.WriteString(w, "Accepted")

}

func coolOff(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn off cooling")
	io.WriteString(w, "Accepted")
}

func heatOn(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn on heat")
	io.WriteString(w, "Accepted")
}

func heatOff(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn off heat")
	io.WriteString(w, "Accepted")
}

func fanOn(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn on fan")
	io.WriteString(w, "Accepted")
}

func fanOff(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn off fan")
	io.WriteString(w, "Accepted")
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to get status")
	io.WriteString(w, "fake status")
}
