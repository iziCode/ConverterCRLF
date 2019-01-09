package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	StartReplaceFormatNEL()

	//ReadFromFile("test_in.log")

}

func StartReplaceFormatNEL() {

	var currentFormat, finalFormat, filePath string

	fmt.Println("Введите какой формат файлов (CRLF, CR, LF или ALL) вы хотите преобразовать:")
	_, err := fmt.Scan(&currentFormat)
	CheckErrors("func StartReplaceFormatNEL()", err)

	fmt.Println("Введите какой формат файлов вы хотите получить на выходе CRLF, CR, или LF:")
	_, err = fmt.Scan(&finalFormat)
	CheckErrors("func StartReplaceFormatNEL()", err)

	currentFormat = strings.ToUpper(currentFormat)

	switch currentFormat {
	case "CRLF":
		ReadFromFile(filePath)

	case "CR":
		fmt.Println("bb")

	case "LF":
		fmt.Println("bb")
	case "ALL":
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

func ReadFromFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64000)

	for {
		n, err := file.Read(data)

		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Println(data[:n])
		CheckFormatNEAL("CRLF", "LF", data[:n])

	}
}
func CheckFormatNEAL(this, that string, b []byte) {

	this = strings.ToUpper(this)
	that = strings.ToUpper(that)

	switch this {
	case "CRLF":
		for i := 0; i < len(b); i++ {
			if b[i] == 13 && b[i+1] == 10 {
				if that == "LF" {
					b = append(b[:i], b[i+1:]...)
					i--
				} else {

				}
			}

		}
		WriteInFile(b)
		fmt.Println(b)

	case "CR":
		fmt.Println("bb")

	case "LF":
		fmt.Println("bb")

	default:
		fmt.Println("Введите один из трех форматов CR, LF, CRLF")
	}

}

//Общая проверка всех ошибок
func CheckErrors(methodName string, err error) {
	if err != nil {
		log.Println(methodName, "get errors:", err)
	}
}
