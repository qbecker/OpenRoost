package database

const (
	InsertSensorDataQuery = "INSERT INTO DAILYDATA(Timestamp, Reading) VALUES(?, ?)"
	InsertCurrent         = "UPDATE Settings set setCurrentTemp = ? WHERE APPNAME = 'ROOST'"
	InsertZip             = "UPDATE Settings SET zipCode = ? WHERE APPNAME = 'ROOST'"
	InsertHome            = "UPDATE Settings SET setTempHome = ? WHERE APPNAME = 'ROOST'"
	InsertAway            = "UPDATE Settings SET setTempAway = ? WHERE APPNAME = 'ROOST'"
	InsertIsHome          = "UPDATE Settings SET isHome = ? WHERE APPNAME = 'ROOST'"
	GetTemp               = "SELECT currentTemp FROM Settings WHERE APPNAME = 'ROOST'"
	GetSetTempQuery       = "SELECT setCurrentTemp FROM Settings WHERE APPNAME = 'ROOST'"
	GetSetTempHomeQuery   = "SELECT setTempHome FROM Settings WHERE APPNAME = 'ROOST'"
	GetSetTempAwayQuery   = "SELECT setTempAway FROM Settings WHERE APPNAME = 'ROOST'"
)
