package yoke

import "../DAO/database"

func createStatus() Status {
	ret := Status{setTemp: database.GetSetTemp(),
		currentTemp: database.GetCurrentTemp(),
		homeTemp:    database.GetTempHome(),
		awayTemp:    database.GetTempAway()}
	return ret

	
}
