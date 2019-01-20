package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const CRLF = "CRLF"
const CR = "CR"
const LF = "LF"
const ALL = "ALL"

func main() {
	//TODO: Add filters to files that do not need to be changed
	//TODO: Add output of all files by formats
	//TODO: Add console input file paths
	StartReplaceFormatNEL()
}

func StartReplaceFormatNEL() {

	currentFormat, finalFormat := ReadInputData()
	inputDataIsCorrected := CheckInputData(currentFormat, finalFormat)

	if inputDataIsCorrected {
		filePaths := GetAllFilesFromCurrentDir()
		ReadFromFilePathsSlice(currentFormat, finalFormat, filePaths)

	} else {
		StartReplaceFormatNEL()
	}

}

func ReadInputData() (currentFormat, finalFormat string) {
	var err error

	fmt.Println("Введите какой формат файлов (CRLF, CR, LF или ALL) вы хотите преобразовать:")
	_, err = fmt.Scan(&currentFormat)
	CheckErrors("func StartReplaceFormatNEL()", err)

	fmt.Println("Введите какой формат файлов вы хотите получить на выходе CRLF, CR, или LF:")
	_, err = fmt.Scan(&finalFormat)
	CheckErrors("func StartReplaceFormatNEL()", err)

	currentFormat = strings.ToUpper(currentFormat)
	finalFormat = strings.ToUpper(finalFormat)

	return
}

func CheckInputData(currentFormat, finalFormat string) (inputDataIsCorrected bool) {
	if currentFormat != CRLF && currentFormat != CR && currentFormat != LF && currentFormat != ALL {
		fmt.Println("Неправильный входной формат!!! Выберите CRLF, CR, LF или ALL")
		inputDataIsCorrected = false
		return
	} else if finalFormat != CRLF && finalFormat != CR && finalFormat != LF {
		fmt.Println("Неправильный выходной формат!!! Выберите CRLF, CR, или LF")
		inputDataIsCorrected = false
		return
	} else if currentFormat == finalFormat {
		fmt.Println("Форматы должны отличаться!!!")
		inputDataIsCorrected = false
		return
	}

	inputDataIsCorrected = true
	return
}

func WriteInFile(filePath string, b []byte) {

	err := ioutil.WriteFile(filePath, b, 0644)
	CheckErrors("func WriteInFile(b []byte)", err)
}

func ReadFromFilePathsSlice(currentFormat, finalFormat string, filePaths []string) {
	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("func ReadFromFilePathsSlice(filePaths []string)", err)
			os.Exit(1)
		}
		CheckErrors("func ReadFromFilePathsSlice(filePaths []string)", err)

		data := make([]byte, 100000)

		for {
			n, err := file.Read(data)

			if err == io.EOF {
				break
			}
			ChangeFormatNEAL(filePath, currentFormat, finalFormat, data[:n])

		}
		err = file.Close()
		CheckErrors("func ReadFromFilePathsSlice(filePaths []string)", err)
	}

}
func ChangeFormatNEAL(filePath, currentFormat, finalFormat string, b []byte) {
	fileByteChanged := false
	if strings.Contains(filePath, `ConverterCRLF.exe`) {
		return
	}
	switch currentFormat {
	case CRLF:
		for i := 0; i < len(b)-1; i++ {
			if b[i] == 13 && b[i+1] == 10 {
				fileByteChanged = true
				if finalFormat == LF {
					b = append(b[:i], b[i+1:]...)
					i--
				} else {
					b = append(b[:i+1], b[i+2:]...)
				}
			}
		}
		if fileByteChanged {
			WriteInFile(filePath, b)
		}

	case CR:
		for i := 0; i < len(b); i++ {
			if b[i] == 13 && i != len(b)-1 && !(b[i+1] == 10) {
				fileByteChanged = true
				if finalFormat == CRLF {
					b = append(b[:i+1], append([]byte{10}, b[i+1:]...)...)
				} else {
					b[i] = 10
				}
			} else if i == len(b)-1 && b[i] == 13 {
				fileByteChanged = true
				if finalFormat == CRLF {
					b = append(b[:], 10)
				} else {
					b[i] = 10
				}

			}
		}
		if fileByteChanged {
			WriteInFile(filePath, b)
		}

	case LF:
		for i := 0; i < len(b); i++ {
			if b[i] == 10 && i > 0 && !(b[i-1] == 13) {
				fileByteChanged = true
				if finalFormat == CRLF {
					b = append(b[:i], append([]byte{13}, b[i:]...)...)
				} else {
					b[i] = 13
				}
			}
		}
		if fileByteChanged {
			WriteInFile(filePath, b)
		}

	case ALL:
		if finalFormat == CRLF {
			ChangeFormatNEAL(filePath, CR, CRLF, b)
			ChangeFormatNEAL(filePath, LF, CRLF, b)
			fileByteChanged = true
		} else if finalFormat == CR {
			ChangeFormatNEAL(filePath, CRLF, CR, b)
			ChangeFormatNEAL(filePath, LF, CR, b)
			fileByteChanged = true
		} else if finalFormat == LF {
			ChangeFormatNEAL(filePath, CRLF, LF, b)
			ChangeFormatNEAL(filePath, CR, LF, b)
			fileByteChanged = true
		}
		if fileByteChanged {
			WriteInFile(filePath, b)
		}

	default:
		fmt.Println("Введите один из трех форматов CR, LF, CRLF")
	}

}

func GetAllFilesFromCurrentDir() (allFilesFromCurrentDir []string) {
	currentPath, err := os.Getwd()
	CheckErrors("func StartReplaceFormatNEL()", err)

	return GetAllFilesFromPath(currentPath)
}

func GetAllFilesFromPath(filePath string) (allFilesFromPath []string) {
	files, err := ioutil.ReadDir(filePath)
	CheckErrors("func GetAllFiles(filePath string)", err)

	for _, file := range files {
		dirOrFilePath := strings.Join([]string{filePath, file.Name()}, "\\")
		if file.IsDir() {
			allFilesFromPath = append(allFilesFromPath, GetAllFilesFromPath(dirOrFilePath)...)
		} else {
			allFilesFromPath = append(allFilesFromPath, dirOrFilePath)
		}
	}
	return
}

//Общая проверка всех ошибок
func CheckErrors(methodName string, err error) {
	if err != nil {
		log.Println(methodName, "get errors:", err)
	}
}
