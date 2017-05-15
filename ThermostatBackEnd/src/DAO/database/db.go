package database

import _ "github.com/mattn/go-sqlite3"
import "fmt"
import "log"
import "database/sql"
import "time"

const (
	dbFile = "roost.db"
	driver = "sqlite3"
)

func InitDB() {
	ExecuteTransactionalDDL(Schema1)
	ExecuteTransactionalDDL(Schema2)
	ExecuteTransactionalDDL(Schema3)
	ExecuteTransactionalDDL(Schema4)
}

func GetSetTemp()int{
	var data int
	transaction, err := getDB().Begin()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := transaction.Query("SELECT setCurrentTemp FROM Settings WHERE APPNAME = 'ROOST'")
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
	rows.Close()
	return data
}

func GetCurrentTemp() int{
	tester:
	var data int
	transaction, err := getDB().Begin()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := transaction.Query(getTemp)
	if err != nil {
		log.Println("it fucked up")
		goto tester
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			log.Println("Err")
		}
	}
	rows.Close()
	return data
}
func InsertHomeOrAway(data int){
	transaction, err := getDB().Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertIsHome)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(data)
	if err != nil {
		log.Fatal(err)
	}
	transaction.Commit()
}
func InsertHomeTemp(data int){
	transaction, err := getDB().Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertHome)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(data)
	if err != nil {
		log.Fatal(err)
	}
	transaction.Commit()
}
func InsertAwayTemp(data int){
	transaction, err := getDB().Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertAway)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(data)
	if err != nil {
		log.Fatal(err)

	}
	transaction.Commit()
}
func InsertZipCode(data int){
	transaction, err := getDB().Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertZip)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(data)
	if err != nil {
		log.Fatal(err)

	}
	transaction.Commit()
}

func InsertCurrentSetTemp(data int) {
	transaction, err := getDB().Begin()
	if err != nil {
		log.Fatal(err)
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

func InsertSensorData(data int) {
	transaction, err := getDB().Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertSensorDataQuery)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	 secs := now.Unix()
	 log.Println(secs)
	_, err = statement.Exec(fmt.Sprintf("%d", secs), data)
	if err != nil {
		log.Fatal(err)
	}
	transaction.Commit()

}
func ExecuteTransactionalSingleRowQuery(query string, selection []interface{}, targets ...interface{}) error {
	transaction, err := getDB().Begin()
	if err != nil {
		return err
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(query)
	if err != nil {
		return err
	}
	row := statement.QueryRow(selection...)
	if err := row.Scan(targets...); err != nil {
		transaction.Rollback()
		return err
	}
	return nil
}

func ExecuteTransactionalDDL(query string, args ...interface{}) error {
	transaction, err := getDB().Begin()
	defer transaction.Commit()
	if err != nil {
		return err
	}
	stmt, err := transaction.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(args...); err != nil {
		transaction.Rollback()
		return err
	}
	return nil
}

var getDB = func() func() *sql.DB {
	db, err := sql.Open(driver, "../foo.db")
	if err != nil {
		panic(err)
	}
	//db.SetMaxOpenConns(1)
	return func() *sql.DB {
		return db
	}
}()
