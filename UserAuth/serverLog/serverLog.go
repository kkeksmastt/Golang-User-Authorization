package serverlog

import (
	c "UserAuth/color"
	"log"
	"os"
)

func Info(str string) {
	log.New(os.Stdout, c.Green+"INFO\t"+c.Reset, log.Ldate|log.Ltime).Printf(str)
}

func ErrorLog(err error) {
	log.New(os.Stderr, c.Red+"ERROR\t"+c.Reset, log.Ldate|log.Ltime|log.Lshortfile).Println(err)
}

func ErrorFatal(err error) {
	log.New(os.Stderr, c.Red+"ERROR\t"+c.Reset, log.Ldate|log.Ltime|log.Lshortfile).Fatal(err)
}
