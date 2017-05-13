package yoke

import (
	"../DAO/database"
	//"encoding/json"
	"io"
	//"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func setTemp(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to set temp from yoke")
	log.Println(r.URL.Path)
	temp := strings.Split(r.URL.Path, "/")
	log.Println(temp[2])
	i, err := strconv.Atoi(temp[2])
	if err != nil {
		// handle error
		log.Println(err)
		io.WriteString(w, "denied")

	} else {
		log.Println(i)
		database.InsertCurrentSetTemp(i)
		io.WriteString(w, "Accepted")
	}

}

func coolOn(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn on cooling from yoke")
	io.WriteString(w, "Accepted")
}
func heatOn(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn on heat from yoke")
	io.WriteString(w, "Accepted")
}
func coolOff(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn off cooling from yoke")
	io.WriteString(w, "Accepted")
}
func heatOff(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn off  heating from yoke")
	io.WriteString(w, "Accepted")
}
func fanOn(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn on fan from yoke")
	io.WriteString(w, "Accepted")
}
func fanOff(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to turn off fan from yoke")
	io.WriteString(w, "Accepted")
}
func getStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to get status from yoke")
	io.WriteString(w, "Accepted")
}
