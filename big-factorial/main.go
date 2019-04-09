package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int32) {
	var p = new(big.Int).SetInt64(1)

	for i := 2; i <= int(n); i++ {
		var p2 = new(big.Int).SetInt64(int64(i))
		p.Mul(p, p2)

	}

	fmt.Println(p)
}

func main() {
	f, err := os.Open("input.txt")
	checkError(err)

	reader := bufio.NewReaderSize(f, 1024*1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	extraLongFactorials(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
