package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the roadsAndLibraries function below.
func roadsAndLibraries(n int32, cLib int32, cRoad int32, cities [][]int32) int64 {
	if cRoad >= cLib {
		return int64(n * cLib)
	}

	var rs int32

	for i := 0; i < int(n); i++ {
		if i == 0 {
			rs += cLib
		}
	}

	return int64(rs)
}

func main() {
	filePath := "input.txt"
	output := "output.txt"

	f, err := os.Open(filePath)
	checkError(err)

	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create(output)
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nmcLibcRoad := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nmcLibcRoad[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nmcLibcRoad[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		cLibTemp, err := strconv.ParseInt(nmcLibcRoad[2], 10, 64)
		checkError(err)
		cLib := int32(cLibTemp)

		cRoadTemp, err := strconv.ParseInt(nmcLibcRoad[3], 10, 64)
		checkError(err)
		cRoad := int32(cRoadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(readLine(reader), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, cLib, cRoad, cities)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
