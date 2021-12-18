package main

import (
	"flag"
	"log"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	Log      bool
}

func main() {
	dbtype := flag.String("db", "mysql", "название директории для кода базы данных")
	Log := flag.Bool("log", true, "Писать логи или нет.")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		infoLog:  infoLog,
		errorLog: errorLog,
		Log:      *Log,
	}

	app.initWeb(dbtype)
	app.initUI()
}
