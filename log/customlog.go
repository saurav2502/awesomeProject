package log

import (
	"log"
	"net/http"
	"os"
)

var (
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func Init() {
	createEmptyFile()
	InitInfoLogger()
	InitErrorLogger()
	InitDebugLogger()
}
func InitInfoLogger() {

	file, err := os.OpenFile("./logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func InitErrorLogger() {
	file, err := os.OpenFile("./logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InitDebugLogger() {
	file, err := os.OpenFile("./logs/debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	DebugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetRequestURI(req *http.Request) string {
	uri := req.RequestURI
	return uri
}

func createEmptyFile() {
	path, _ := os.Getwd()
	if _, err := os.Stat(path + "/logs"); os.IsNotExist(err) {
		err := os.Mkdir(path+"/logs", 0700)
		if err != nil {
			ErrorLogger.Printf("Error while creating logs directory %v", err)
		}
	}
}
