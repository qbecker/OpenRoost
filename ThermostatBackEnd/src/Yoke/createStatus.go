package yoke

import "../DAO/database"

func createStatus() Status {
	ret := Status{SetTemp: database.GetSetTemp(),
		CurrentTemp: database.GetCurrentTemp(),
		HomeTemp:    database.GetTempHome(),
		AwayTemp:    database.GetTempAway()}
	return ret

}
