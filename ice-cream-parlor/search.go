package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

// Complete the whatFlavors function below.
func whatFlavors(cost []int, money int) {

	costref := make([]int, 0, len(cost))

	for _, i := range cost {
		costref = append(costref, i)
	}

	sort.Ints(cost)
	var index int
	index = len(cost) - 1

	if money <= cost[len(cost)-1] {
		index = findIndex(cost, money, 0, len(cost)-1)
	}

	cost2 := cost[:index+1]

	cost3 := 0
	cost4 := 0

	for i := 0; i < len(cost2); i++ {
		e := money - cost2[i]

		index = findIndex(cost2, e, 0, len(cost2)-1)
		if index >= 0 {
			cost3 = cost2[i]
			cost4 = e
		}
	}

	pp.Println(cost3, cost4)
	for i := 0; i < len(costref); i++ {
		if costref[i] == cost3 {
			fmt.Print(i+1, " ")
			continue
		}

		if costref[i] == cost4 {
			fmt.Print(i + 1)
		}
	}

	fmt.Println()

}

func findIndex(ac []int, in int, l, h int) int {
	var m int

	for l < h {
		m = (l + h) / 2
		if in > ac[m] {
			l = m + 1
			continue
		}

		h = m
	}

	if ac[l] == in {
		return l
	}

	return -1
}

func main() {
	filePath := "/Users/namnh/workspace/go/src/github.com/namnhce/my-go-practices/search/input.txt"
	f, err := os.Open(filePath)
	checkError(err)

	reader := bufio.NewReaderSize(f, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		money := int(moneyTemp)

		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int(nTemp)

		costTemp := strings.Split(readLine(reader), " ")

		var cost []int

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
	}
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
