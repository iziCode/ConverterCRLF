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

	PrintAllFilesAndFormats(allFilesAndFormats)
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

func getFileFormatNEL(allFilesPaths []string) (allFilesAndFormats []map[string]string) {
	mapCRLF := make(map[string]string)
	mapCR := make(map[string]string)
	mapLF := make(map[string]string)
	allFilesAndFormats = append(allFilesAndFormats, mapCRLF)
	allFilesAndFormats = append(allFilesAndFormats, mapCR)
	allFilesAndFormats = append(allFilesAndFormats, mapLF)

	funcName := "getFileFormatNEL(allFilesPaths []string) (allFilesAndFormats map[string]string)"
	for _, filePath := range allFilesPaths {
		dataBytes, err := ioutil.ReadFile(filePath)
		CheckErrors(funcName, err)

		for i := 0; i < len(dataBytes); i++ {

			if dataBytes[i] == 13 && i+1 < len(dataBytes) && dataBytes[i+1] == 10 {
				allFilesAndFormats[0][filePath] = CRLF
				break
			} else if dataBytes[i] == 13 {
				allFilesAndFormats[1][filePath] = CR
				break
			} else if dataBytes[i] == 10 {
				allFilesAndFormats[2][filePath] = LF
				break
			}
		}

	}
	return allFilesAndFormats
}

func PrintAllFilesAndFormats(allFilesAndFormats []map[string]string) {

	for _, fileFormatMap := range allFilesAndFormats {
		for file, format := range fileFormatMap {
			fmt.Println("Формат:", format, "Файл:", file)
		}
	}
}
