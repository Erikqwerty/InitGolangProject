package main

import (
	"os"
)

func (app application) initWeb(dbtype *string) {
	err := os.MkdirAll("./cmd/web/", 0777)
	if err != nil {
		app.errorLog.Println(err.Error())
	} else if app.Log {
		app.infoLog.Println("Инициализированны директории ./cmd/web")
	}
	err = os.MkdirAll("./pkg/models/"+*dbtype, 0777)
	if err != nil {
		app.errorLog.Println(err.Error())
	} else if app.Log {
		app.infoLog.Println("Инициализированны директории ./pkg/models")
	}
	f, err := os.Create("./cmd/web/main.go")
	if err != nil {
		app.errorLog.Println(err.Error())
	} else if app.Log {
		app.infoLog.Println("Создан файл ./cmd/web/main.go ")
	}
	defer f.Close()

	_, err = f.WriteString("package main \n\nfunc main(){\n\n}")

	if err != nil {
		app.errorLog.Println(err.Error())
	} else if app.Log {
		app.infoLog.Println("Инициализированна функция main() в файле main.go")
	}
}

func (app application) initUI() {
	err := os.MkdirAll("./ui/html", 0777)
	if err != nil {
		app.errorLog.Println(err.Error())
	} else if app.Log {
		app.infoLog.Println("Инициализированны директории ./ui/html")
	}
	err = os.MkdirAll("./ui/static/css", 0777)
	if err != nil {
		app.errorLog.Println(err.Error())
	} else if app.Log {
		app.infoLog.Println("Инициализированны директории ./ui/static/css")
	}
	err = os.MkdirAll("./ui/static/js", 0777)
	if err != nil {
		app.errorLog.Println(err.Error())
	} else if app.Log {
		app.infoLog.Println("Инициализированны директории ./ui/static/js")
	}
	err = os.MkdirAll("./ui/static/img", 0777)
	if err != nil {
		app.errorLog.Println(err.Error())
	} else if app.Log {
		app.infoLog.Println("Инициализированны директории ./ui/static/img")
	}
}
