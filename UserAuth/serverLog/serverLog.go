package serverlog

import (
	"log"
	"os"
)

func Info(str string) {
	log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime).Printf(str)
}

func ErrorLog(err error) {
	log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile).Println(err)
}

func ErrorFatal(err error) {
	log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile).Fatal(err)
}
