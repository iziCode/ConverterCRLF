package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func ShowAllFilesAndFormats() {

	filePath := ReadShowAllFilesAndFormats()

	allFilesPaths := GetAllFilesFromPath(filePath)

	allFilesAndFormats := getFileFormatNEL(allFilesPaths)

	fmt.Println(allFilesAndFormats)
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
	funcName := "getFileFormatNEL(allFilesPaths []string) (allFilesAndFormats map[string]string)"
	allFilesAndFormats = make(map[string]string)
	for _, filePath := range allFilesPaths {
		dataBytes, err := ioutil.ReadFile(filePath)
		CheckErrors(funcName, err)

		for i := 0; i < len(dataBytes); i++ {

			if dataBytes[i] == 13 && i+1 < len(dataBytes) && dataBytes[i+1] == 10 {
				allFilesAndFormats[filePath] = CRLF
				break
			} else if dataBytes[i] == 13 {
				allFilesAndFormats[filePath] = CR
				break
			} else if dataBytes[i] == 10 {
				allFilesAndFormats[filePath] = LF
				break
			}
		}

	}
	return allFilesAndFormats
}
