package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func exercise2(in string) string {
	rs := ""
	for _, rune := range in {
		s := string(rune)
		if s == "E" {
			rs += "S"
		} else {
			rs += "E"
		}
	}
	return rs
}

func main() {
	filePath := "input.txt"

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReaderSize(f, 1024*1024)
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		panic(err)
	}

	q := int32(qTemp)

	for i := 0; i < int(q); i++ {

		_, err = strconv.ParseInt(readLine(reader), 10, 64)
		input := readLine(reader)

		rs := exercise2(input)
		fmt.Printf("Case #%v: %v\n", i+1, rs)

	}

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
