package yoke

import (
	"../DAO/database"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func isHome(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to set isHome from yoke")
	log.Println(r.URL.Path)
	temp := strings.Split(r.URL.Path, "/")
	log.Println(temp[2])
	i, err := strconv.Atoi(temp[2])
	if err != nil {
		// handle error
		log.Println(err)
		io.WriteString(w, "denied")

	} else {
		if i > 1 {
			io.WriteString(w, "denied")
		} else {
			log.Println(i)
			database.InsertHomeOrAway(i)
			io.WriteString(w, "Accepted")
		}
	}
}
func isAway(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to set isAway from yoke")
	log.Println(r.URL.Path)
	temp := strings.Split(r.URL.Path, "/")
	log.Println(temp[2])
	i, err := strconv.Atoi(temp[2])
	if err != nil {
		log.Println(err)
		io.WriteString(w, "denied")

	} else {
		if i > 1 {
			io.WriteString(w, "denied")
		} else {
			log.Println(i)
			database.InsertHomeOrAway(i)
			io.WriteString(w, "Accepted")
		}
	}
}

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

func setZipCode(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to set zip code from yoke")
	log.Println(r.URL.Path)
	temp := strings.Split(r.URL.Path, "/")
	log.Println(temp[2])
	i, err := strconv.Atoi(temp[2])
	if err != nil {
		log.Println(err)
		io.WriteString(w, "denied")

	} else {
		log.Println(i)
		database.InsertZipCode(i)
		io.WriteString(w, "Accepted")
	}
}

func setHomeTemp(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to set home temp from yoke")
	log.Println(r.URL.Path)
	temp := strings.Split(r.URL.Path, "/")
	log.Println(temp[2])
	i, err := strconv.Atoi(temp[2])
	if err != nil {
		log.Println(err)
		io.WriteString(w, "denied")

	} else {
		log.Println(i)
		database.InsertHomeTemp(i)
		io.WriteString(w, "Accepted")
	}
}

func setAwayTemp(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to set away temp from yoke")
	log.Println(r.URL.Path)
	temp := strings.Split(r.URL.Path, "/")
	log.Println(temp[2])
	i, err := strconv.Atoi(temp[2])
	if err != nil {
		log.Println(err)
		io.WriteString(w, "denied")

	} else {
		log.Println(i)
		database.InsertAwayTemp(i)
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
	status := createStatus()
	log.Println(status)
	log.Println(status.setTemp)
	//Next step actually return as JSON
	io.WriteString(w, "Accepted")
}
