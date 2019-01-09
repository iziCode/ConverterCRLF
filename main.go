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
	filePaths := GetAllFilesFromCurrentDir()

	fmt.Println("Введите какой формат файлов (CRLF, CR, LF или ALL) вы хотите преобразовать:")
	_, err = fmt.Scan(&currentFormat)
	CheckErrors("func StartReplaceFormatNEL()", err)

	fmt.Println("Введите какой формат файлов вы хотите получить на выходе CRLF, CR, или LF:")
	_, err = fmt.Scan(&finalFormat)
	CheckErrors("func StartReplaceFormatNEL()", err)

	currentFormat = strings.ToUpper(currentFormat)

	switch currentFormat {
	case CRLF:
		ReadFromFilePathsSlice(currentFormat, finalFormat, filePaths)

	case CR:
		fmt.Println("bb")

	case LF:
		fmt.Println("bb")
	case ALL:
		fmt.Println("bb")

	default:
		fmt.Println("Введите корректные данные!!!")
		StartReplaceFormatNEL()
	}

}

func WriteInFile(b []byte) {
	//data := []byte("Hello Bold!")
	file, err := os.Create("test_out.log")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(b)

	fmt.Println("Done.")
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
			fmt.Println(data[:n])
			ChangeFormatNEAL(currentFormat, finalFormat, data[:n])

		}
		err = file.Close()
		CheckErrors("func ReadFromFilePathsSlice(filePaths []string)", err)
	}

}
func ChangeFormatNEAL(currentFormat, finalFormat string, b []byte) {
	currentFormat = strings.ToUpper(currentFormat)
	finalFormat = strings.ToUpper(finalFormat)

	switch currentFormat {
	case CRLF:
		for i := 0; i < len(b); i++ {
			if b[i] == 13 && b[i+1] == 10 {
				if finalFormat == LF {
					b = append(b[:i], b[i+1:]...)
					i--
				} else {

				}
			}

		}
		WriteInFile(b)
		fmt.Println(b)

	case CR:
		fmt.Println("bb")

	case LF:
		fmt.Println("bb")

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
