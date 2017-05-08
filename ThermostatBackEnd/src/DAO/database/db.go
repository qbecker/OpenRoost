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
	//strconv.FormatInt(int64(123), 10)
	_, err = statement.Exec(fmt.Sprintf("%d", time.Now().Unix), data)
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
	return func() *sql.DB {
		return db
	}
}()
