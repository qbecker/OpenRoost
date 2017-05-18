package logger

import(
  "../DAO/database"

)
func WriteLog(message string){
  database.InsertLogMessage(message)
}
