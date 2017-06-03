package logger


//simple interface for writing message to database log
import(
  "../DAO/database"

)
func WriteLog(message string){
  database.InsertLogMessage(message)
}
