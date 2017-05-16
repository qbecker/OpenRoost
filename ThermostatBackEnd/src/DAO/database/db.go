package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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

func setupDB(db *sql.DB) {
	_, err := db.Exec(Schema1)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DatabaseObject) GetTempHome() int {
	var data int
	d.Lock()
	defer d.Unlock()
	rows, err := d.db.Query(GetSetTempHomeQuery)
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

func (d *DatabaseObject) GetTempAway() int {
	var data int
	d.Lock()
	defer d.Unlock()
	rows, err := d.db.Query(GetSetTempAwayQuery)
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

func (d *DatabaseObject) GetSetTemp() int {
	var data int
	d.Lock()
	defer d.Unlock()
	rows, err := d.db.Query(GetSetTempQuery)
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
	rows, err := d.db.Query(GetTemp)
	if err != nil {
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
	statement, err := transaction.Prepare(InsertHomeQuery)
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
	statement, err := transaction.Prepare(InsertZipQuery)
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
