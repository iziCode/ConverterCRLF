package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

const (
	CRLF = "CRLF"
	CR   = "CR"
	LF   = "LF"
	ALL  = "ALL"
)

func main() {
	//TODO: Add filters to files that do not need to be changed
	//TODO: Add output of all files by formats
	StartReplaceFormatNEL()
}

func StartReplaceFormatNEL() {
	funcName := "StartReplaceFormatNEL()"

	filePath, currentFormat, finalFormat := ReadInputData()
	inputDataIsCorrected := CheckInputData(filePath, currentFormat, finalFormat)

	if inputDataIsCorrected {
		filePaths := GetAllFilesFromPath(filePath)

		for _, filePath := range filePaths {
			dataBytes, err := ioutil.ReadFile(filePath)
			CheckErrors(funcName, err)
			ChangeFormatNEAL(filePath, currentFormat, finalFormat, dataBytes)
		}

	} else {
		StartReplaceFormatNEL()
	}

}

func ReadInputData() (filePath, currentFormat, finalFormat string) {
	var err error
	funcName := "ReadInputData() (filePath, currentFormat, finalFormat string)"

	fmt.Println("Вставьте путь до нужной папки:")
	_, err = fmt.Scan(&filePath)
	CheckErrors("func StartReplaceFormatNEL()", err)

	fmt.Println("Введите какой формат файлов (CRLF, CR, LF или ALL) вы хотите преобразовать:")
	_, err = fmt.Scan(&currentFormat)
	CheckErrors("func StartReplaceFormatNEL()", err)

	fmt.Println("Введите какой формат файлов вы хотите получить на выходе CRLF, CR, или LF:")
	_, err = fmt.Scan(&finalFormat)
	CheckErrors("func StartReplaceFormatNEL()", err)

	currentFormat = strings.ToUpper(currentFormat)
	finalFormat = strings.ToUpper(finalFormat)

	log.WithFields(log.Fields{
		"filePath":      filePath,
		"currentFormat": currentFormat,
		"finalFormat":   finalFormat,
	}).Info(funcName)

	return filePath, currentFormat, finalFormat
}

func CheckInputData(filePath, currentFormat, finalFormat string) (inputDataIsCorrected bool) {
	funcName := "CheckInputData(filePath, currentFormat, finalFormat string) (inputDataIsCorrected bool)"

	log.WithFields(log.Fields{
		"filePath":      filePath,
		"currentFormat": currentFormat,
		"finalFormat":   finalFormat,
	}).Info(funcName)

	if filePath == "" {
		fmt.Println("Путь до папки не может быть пустым!!!")
		return false
	}
	if currentFormat != CRLF && currentFormat != CR && currentFormat != LF && currentFormat != ALL {
		fmt.Println("Неправильный входной формат!!! Выберите CRLF, CR, LF или ALL")
		return false
	} else if finalFormat != CRLF && finalFormat != CR && finalFormat != LF {
		fmt.Println("Неправильный выходной формат!!! Выберите CRLF, CR, или LF")
		return false
	} else if currentFormat == finalFormat {
		fmt.Println("Форматы должны отличаться!!!")
		return false
	}

	return true
}

func WriteInFile(filePath string, dataBytes []byte) {
	funcName := "WriteInFile(filePath string, dataBytes []byte)"

	log.WithFields(log.Fields{
		"filePath":  filePath,
		"dataBytes": dataBytes,
	}).Info(funcName)

	err := ioutil.WriteFile(filePath, dataBytes, 0644)
	CheckErrors(funcName, err)
}

func ChangeFormatNEAL(filePath, currentFormat, finalFormat string, dataBytes []byte) {
	fileByteChanged := false
	funcName := "ChangeFormatNEAL(filePath, currentFormat, finalFormat string, dataBytes []byte)"

	log.WithFields(log.Fields{
		"filePath":      filePath,
		"currentFormat": currentFormat,
		"finalFormat":   finalFormat,
		"dataBytes":     dataBytes,
	}).Info(funcName)

	if strings.Contains(filePath, `ConverterCRLF.exe`) {
		return
	}
	switch currentFormat {
	case CRLF:
		for i := 0; i < len(dataBytes)-1; i++ {
			if dataBytes[i] == 13 && dataBytes[i+1] == 10 {
				fileByteChanged = true
				if finalFormat == LF {
					dataBytes = append(dataBytes[:i], dataBytes[i+1:]...)
					i--
				} else {
					dataBytes = append(dataBytes[:i+1], dataBytes[i+2:]...)
				}
			}
		}
		if fileByteChanged {
			WriteInFile(filePath, dataBytes)
		}

	case CR:
		for i := 0; i < len(dataBytes); i++ {
			if dataBytes[i] == 13 && i != len(dataBytes)-1 && !(dataBytes[i+1] == 10) {
				fileByteChanged = true
				if finalFormat == CRLF {
					dataBytes = append(dataBytes[:i+1], append([]byte{10}, dataBytes[i+1:]...)...)
				} else {
					dataBytes[i] = 10
				}
			} else if i == len(dataBytes)-1 && dataBytes[i] == 13 {
				fileByteChanged = true
				if finalFormat == CRLF {
					dataBytes = append(dataBytes[:], 10)
				} else {
					dataBytes[i] = 10
				}

			}
		}
		if fileByteChanged {
			WriteInFile(filePath, dataBytes)
		}

	case LF:
		for i := 0; i < len(dataBytes); i++ {
			if dataBytes[i] == 10 && i > 0 && !(dataBytes[i-1] == 13) {
				fileByteChanged = true
				if finalFormat == CRLF {
					dataBytes = append(dataBytes[:i], append([]byte{13}, dataBytes[i:]...)...)
				} else {
					dataBytes[i] = 13
				}
			}
		}
		if fileByteChanged {
			WriteInFile(filePath, dataBytes)
		}

	case ALL:
		if finalFormat == CRLF {
			ChangeFormatNEAL(filePath, CR, CRLF, dataBytes)
			ChangeFormatNEAL(filePath, LF, CRLF, dataBytes)
			fileByteChanged = true
		} else if finalFormat == CR {
			ChangeFormatNEAL(filePath, CRLF, CR, dataBytes)
			ChangeFormatNEAL(filePath, LF, CR, dataBytes)
			fileByteChanged = true
		} else if finalFormat == LF {
			ChangeFormatNEAL(filePath, CRLF, LF, dataBytes)
			ChangeFormatNEAL(filePath, CR, LF, dataBytes)
			fileByteChanged = true
		}
		if fileByteChanged {
			WriteInFile(filePath, dataBytes)
		}

	default:
		fmt.Println("Введите один из трех форматов CR, LF, CRLF")
	}

}

func GetAllFilesFromPath(filePath string) (allFilesFromPath []string) {
	funcName := "GetAllFilesFromPath(filePath string) (allFilesFromPath []string)"

	files, err := ioutil.ReadDir(filePath)
	CheckErrors(funcName, err)

	for _, file := range files {
		dirOrFilePath := strings.Join([]string{filePath, file.Name()}, "\\")
		if file.IsDir() {
			allFilesFromPath = append(allFilesFromPath, GetAllFilesFromPath(dirOrFilePath)...)
		} else {
			allFilesFromPath = append(allFilesFromPath, dirOrFilePath)
		}
	}

	log.WithFields(log.Fields{
		"filePath":         filePath,
		"allFilesFromPath": allFilesFromPath,
	}).Info(funcName)

	return allFilesFromPath
}

func CheckErrors(funcName string, err error) {
	if err != nil {
		log.Errorf(funcName, "get errors:", err)
	}
}
