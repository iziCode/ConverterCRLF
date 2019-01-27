package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func ShowAllFilesAndFormats() {

	filePath := ReadShowAllFilesAndFormats()

	allFilesPaths := GetAllFilesFromPath(filePath)

	getFileFormatNEL(allFilesPaths)

}

func ReadShowAllFilesAndFormats() (filePath string) {
	var err error
	funcName := "ReadShowAllFilesAndFormats() (filePath, currentFormat, finalFormat string)"

	fmt.Println("Вставьте путь до нужной папки:")
	_, err = fmt.Scan(&filePath)
	CheckErrors(funcName, err)

	log.WithFields(log.Fields{
		"filePath": filePath,
	}).Info(funcName)

	return filePath
}

func getFileFormatNEL(allFilesPaths []string) (allFilesAndFormats map[string]string) {

	return
}
