package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	p1 = [][]int{
		{4, 9, 2},
		{3, 5, 7},
		{8, 1, 6}}

	p2 = [][]int{
		{8, 1, 6},
		{3, 5, 7},
		{4, 9, 2}}

	p3 = [][]int{
		{8, 3, 4},
		{1, 5, 9},
		{6, 7, 2}}

	p4 = [][]int{
		{6, 7, 2},
		{1, 5, 9},
		{8, 3, 4}}
)

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int) int {
	var s1, s2, s3, s4, s5, s6, s7, s8 float64
	n := len(s)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s1 += math.Abs(float64(s[i][j] - p1[i][j]))
			s2 += math.Abs(float64(s[i][j] - p1[i][2-j]))
			s3 += math.Abs(float64(s[i][j] - p2[i][j]))
			s4 += math.Abs(float64(s[i][j] - p2[i][2-j]))
			s5 += math.Abs(float64(s[i][j] - p3[i][j]))
			s6 += math.Abs(float64(s[i][j] - p3[i][2-j]))
			s7 += math.Abs(float64(s[i][j] - p4[i][j]))
			s8 += math.Abs(float64(s[i][j] - p4[i][2-j]))
		}
	}
	rs := min(min(s1, s2), min(s3, s4), min(s5, s6), min(s7, s8))
	return int(rs)
}

func min(a ...float64) float64 {
	min := a[0]

	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	filePath := "/Users/namnh/workspace/go/src/github.com/namnhce/my-go-practices/magic-square/input.txt"
	output := "/Users/namnh/workspace/go/src/github.com/namnhce/my-go-practices/magic-square/output.txt"

	f, err := os.Open(filePath)
	checkError(err)

	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create(output)
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var s [][]int
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(readLine(reader), " ")

		var sRow []int
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
