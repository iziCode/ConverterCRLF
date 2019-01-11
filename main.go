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
	StartReplaceFormatNEL()
}

func StartReplaceFormatNEL() {
	var err error
	var currentFormat, finalFormat string

	fmt.Println("Введите какой формат файлов (CRLF, CR, LF или ALL) вы хотите преобразовать:")
	_, err = fmt.Scan(&currentFormat)
	CheckErrors("func StartReplaceFormatNEL()", err)

	fmt.Println("Введите какой формат файлов вы хотите получить на выходе CRLF, CR, или LF:")
	_, err = fmt.Scan(&finalFormat)
	CheckErrors("func StartReplaceFormatNEL()", err)

	currentFormat = strings.ToUpper(currentFormat)
	finalFormat = strings.ToUpper(finalFormat)

	if currentFormat != CRLF && currentFormat != CR && currentFormat != LF && currentFormat != ALL {
		fmt.Println("Неправильный входной формат!!! Выберите CRLF, CR, LF или ALL")
		StartReplaceFormatNEL()
	} else if finalFormat != CRLF && finalFormat != CR && finalFormat != LF {
		fmt.Println("Неправильный выходной формат!!! Выберите CRLF, CR, или LF")
		StartReplaceFormatNEL()
	} else if currentFormat == finalFormat {
		fmt.Println("Форматы должны отличаться!!!")
		StartReplaceFormatNEL()
	} else {
		filePaths := GetAllFilesFromCurrentDir()
		ReadFromFilePathsSlice(currentFormat, finalFormat, filePaths)
	}

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

			if err == io.EOF { // если конец файла
				break // выходим из цикла
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

	default:
		fmt.Println("Введите один из трех форматов CR, LF, CRLF")
	}

}

func GetAllFilesFromCurrentDir() (allFilesFromCurrentDir []string) {
	//currentPath, err := os.Getwd()
	//CheckErrors("func StartReplaceFormatNEL()", err)
	currentPath := `D:\VM\image-docker\ci-testing — копия`

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
