package database

var db = InitDB()

func GetSetTemp()int{
    return db.GetSetTemp()
}

func GetCurrentTemp()int{
    return db.GetCurrentTemp()
}

func InsertHomeOrAway(data int){
    db.InsertHomeOrAway(data)
}

func InsertHomeTemp(data int){
    db.InsertHomeTemp(data)
}

func InsertAwayTemp(data int){
    db.InsertAwayTemp(data)
}

func InsertZipCode(data int){
    db.InsertZipCode(data)
}

func InsertCurrentSetTemp(data int){
    db.InsertCurrentSetTemp(data)
}

func InsertSensorData(data int){
    db.InsertSensorData(data)
}