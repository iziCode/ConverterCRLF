package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	ReadFromFile("test_in.log")

}

func ReplaceFormatNEL(this, that, filePath string) {

	this = strings.ToUpper(this)

	switch this {
	case "CRLF":
		ReadFromFile(filePath)

	case "CR":
		fmt.Println("bb")

	case "LF":
		fmt.Println("bb")

	default:
		fmt.Println("Введите один из трех форматов CR, LF, CRLF")
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
		CheckFortmatNEL("CRLF", "LF", data[:n])

	}
}
func CheckFortmatNEL(this, that string, b []byte) {

	this = strings.ToUpper(this)
	that = strings.ToUpper(that)

	switch this {
	case "CRLF":
		for i := 0; i< len(b); i++ {
			if b[i] == 13 && b[i+1] == 10 {
				if that == "LF" {
					b = append(b[:i] , b[i+1:]...)
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
