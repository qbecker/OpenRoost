package database

import (
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

type DatabaseObject struct {
	db *sql.DB
	sync.RWMutex
}

const (
	dbFile = "roost.db"
	driver = "sqlite3"
)

func InitDB() *DatabaseObject {
	log.Println("starting db in db.go")
	db, err := sql.Open("sqlite3", "database_file.sqlite")
	if err != nil {
		log.Fatal("could not open sqlite3 database file", err)
	}
	//defer db.Close()
	setupDB(db)
	return &DatabaseObject{db, sync.RWMutex{}}
}

func (d *DatabaseObject) GetSetTemp() int {
	var data int
	d.Lock()
	defer d.Unlock()
	rows, err := d.db.Query("SELECT setCurrentTemp FROM Settings WHERE APPNAME = 'ROOST'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			log.Println("Err")
		}
	}
	return data
}

func (d *DatabaseObject) GetCurrentTemp() int {
tester:
	var data int
	d.Lock()
	defer d.Unlock()
	rows, err := d.db.Query(getTemp)
	if err != nil {
		log.Println("it fucked up")
		log.Println(err)
		goto tester
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			log.Println("Err")
		}
	}
	return data
}
func (d *DatabaseObject) InsertHomeOrAway(data int) {
	d.Lock()
	defer d.Unlock()
	transaction, err := d.db.Begin()
	if err != nil {
		fmt.Printf("begin. Exec error=%s", err)
		return
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertSensorDataQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(data)
	if err != nil {
		log.Fatal(err)

	}
	transaction.Commit()
}
func (d *DatabaseObject) InsertHomeTemp(data int) {
	d.Lock()
	defer d.Unlock()
	transaction, err := d.db.Begin()
	if err != nil {
		fmt.Printf("begin. Exec error=%s", err)
		return
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertSensorDataQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(data)
	if err != nil {
		log.Fatal(err)

	}
	transaction.Commit()
}
func (d *DatabaseObject) InsertAwayTemp(data int) {
	d.Lock()
	defer d.Unlock()
	transaction, err := d.db.Begin()
	if err != nil {
		fmt.Printf("begin. Exec error=%s", err)
		return
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertSensorDataQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(data)
	if err != nil {
		log.Fatal(err)

	}
	transaction.Commit()
}
func (d *DatabaseObject) InsertZipCode(data int) {
	d.Lock()
	defer d.Unlock()
	transaction, err := d.db.Begin()
	if err != nil {
		fmt.Printf("begin. Exec error=%s", err)
		return
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertSensorDataQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(data)
	if err != nil {
		log.Fatal(err)

	}
	transaction.Commit()
}

func (d *DatabaseObject) InsertCurrentSetTemp(data int) {
	d.Lock()
	defer d.Unlock()
	transaction, err := d.db.Begin()
	if err != nil {
		fmt.Printf("begin. Exec error=%s", err)
		return
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertCurrent)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(data)
	if err != nil {
		log.Fatal(err)

	}
	transaction.Commit()
}

func (d *DatabaseObject) InsertSensorData(data int) {
	log.Print("Inserting Data ")
	log.Println(data)
	d.Lock()
	defer d.Unlock()
	now := time.Now()
	secs := now.Unix()
	transaction, err := d.db.Begin()
	if err != nil {
		fmt.Printf("begin. Exec error=%s", err)
		return
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertSensorDataQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(fmt.Sprintf("%d", secs), data)
	if err != nil {
		log.Fatal(err)

	}
	transaction.Commit()
}

func setupDB(db *sql.DB) {
	_, err := db.Exec(Schema1)
	if err != nil {
		log.Fatal(err)
	}
}
