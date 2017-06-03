package database


//interface to database, keeping the options for diffrent database systems open
var db = InitDB()

func GetSetTemp() int {
	return db.GetSetTemp()
}

func GetTempHome() int {
	return db.GetTempHome()
}

func GetTempAway() int {
	return db.GetTempAway()
}
func GetCurrentTemp() int {
	return db.GetCurrentTemp()
}

func InsertHomeOrAway(data int) {
	db.InsertHomeOrAway(data)
}

func InsertHomeTemp(data int) {
	db.InsertHomeTemp(data)
}

func InsertAwayTemp(data int) {
	db.InsertAwayTemp(data)
}

func InsertZipCode(data int) {
	db.InsertZipCode(data)
}

func InsertCurrentSetTemp(data int) {
	db.InsertCurrentSetTemp(data)
}

func InsertSensorData(data int) {
	db.InsertSensorData(data)
}

func InsertLogMessage(data string){
	db.InsertLogMessage(data)
}
