package database

const (
	InsertSensorDataQuery = "INSERT INTO DAILYDATA(Timestamp, Reading) VALUES(?, ?)"
	InsertCurrent         = "UPDATE Settings set setCurrentTemp = ?"
)
