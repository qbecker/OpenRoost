package database

import _ "github.com/mattn/go-sqlite3"
import "time"

const (
	dbFile = "roost.db"
	driver = "sqlite3"
)

func InitDB() {
	ExecuteTransactionalDDL(Schema)
}

func InsertSensorData(data int) {
	transaction, err := getDB().Begin()
	if err != nil {
		return err
	}
	defer transaction.Commit()
	statement, err := transaction.Prepare(InsertSensorDataQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(data, fmt.Sprintf(time.Now()))
	if err != nil {
		log.Fatal(err)

	}

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

var getDB = func() func() *sql.DB {
	db, err := sql.Open(driver, fmt.Sprintf("%v%v", utils.DatabaseDirectory(), dbFile))
	if err != nil {
		panic(err)
	}
	return func() *sql.DB {
		return db
	}
}()
