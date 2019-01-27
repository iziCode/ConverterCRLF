package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

const (
	CRLF = "CRLF"
	CR   = "CR"
	LF   = "LF"
	ALL  = "ALL"
)

func main() {
	ShowAllFilesAndFormats()
	return
	ChooseMethod()
	//TODO: Add filters to files that do not need to be changed
	//TODO: Add output of all files by formats

}

func ChooseMethod() {
	var err error
	funcName := "ChooseMethod()"
	var ChosenMethod string

	fmt.Println(`Для отображение всех файлов и их форматов напишите "1"`)
	fmt.Println(`Для преобразования файлов в нужный формат напишите "2":`)
	_, err = fmt.Scan(&ChosenMethod)
	CheckErrors(funcName, err)
	CheckChooseMethod(ChosenMethod)

	switch ChosenMethod {
	case "1":
		ShowAllFilesAndFormats()
	case "2":
		StartReplaceFormatNEL()

	}

}

func CheckChooseMethod(ChosenMethod string) {
	if ChosenMethod != "1" && ChosenMethod != "2" {
		fmt.Println(`Неверная команда!`)
		ChooseMethod()
	}
}

func CheckErrors(funcName string, err error) {
	if err != nil {
		log.Errorf(funcName, "get errors:", err)
	}
}
