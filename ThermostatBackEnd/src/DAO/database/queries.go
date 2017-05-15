package database

const (
	InsertSensorDataQuery = "INSERT INTO DAILYDATA(Timestamp, Reading) VALUES(?, ?)"
	InsertCurrent         = "UPDATE Settings set setCurrentTemp = ? WHERE APPNAME = 'ROOST'"
	InsertZip = "UPDATE Settings SET zipCode = ? WHERE APPNAME = 'ROOST'"
	InsertHome = "UPDATE Settings SET setTempHome = ? WHERE APPNAME = 'ROOST'"
	InsertAway = "UPDATE Settings SET setTempAway = ? WHERE APPNAME = 'ROOST'"
	InsertIsHome = "UPDATE Settings SET isHome = ? WHERE APPNAME = 'ROOST'"
	getTemp = "SELECT currentTemp FROM Settings WHERE APPNAME = 'ROOST'"
)
