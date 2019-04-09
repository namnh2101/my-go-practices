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

// Complete the triplets function below.
func triplets(a1 []int, b1 []int, c1 []int) int {
	a := removeDuplicates(a1)
	b := removeDuplicates(b1)
	c := removeDuplicates(c1)

	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)

	rs := 0

	if a[0] > b[len(b)-1] && c[0] > b[len(b)-1] {
		return 0
	}

	for i := 0; i < len(a); i++ {

		pp.Println(fmt.Sprintf("a[%v]=%v", i, a[i]))
		qIndex := findIndex(b, a[i], 0, len(b)-1)
		if qIndex == len(b) {
			continue
		}
		if qIndex < 0 {
			qIndex = 0
		}

		dupB := b[qIndex:]
		pp.Println(fmt.Sprintf("dupB=%v", dupB))
		pp.Println(fmt.Sprintf("len(dupB) = %v", len(dupB)))

		for j := 0; j < len(dupB); j++ {
			pp.Println(fmt.Sprintf("c=%v", c))
			pp.Println(fmt.Sprintf("dupB[%v]=%v", j, dupB[j]))
			if a[i] > dupB[j] {
				continue
			}
			rIndex := findIndex(c, dupB[j], 0, len(c)-1)
			if rIndex < 0 {
				continue
			}
			pp.Println("rIndex = ", rIndex)
			if rIndex == len(c) {
				rIndex = len(c) - 1
			}

			dupC := c[:rIndex+1]

			pp.Println(fmt.Sprintf("dupC=%v", dupC))
			pp.Println(fmt.Sprintf("len(dupC) = %v", len(dupC)))

			rs += len(dupC)
		}
		pp.Println("END")
		pp.Println()
	}
	fmt.Println(rs)
	return rs
}

func triplets2(a1 []int, b1 []int, c1 []int) int {
	a := removeDuplicates(a1)
	b := removeDuplicates(b1)
	c := removeDuplicates(c1)

	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)
	pp.Println(a)
	pp.Println(b)
	pp.Println(c)
	if a[0] > b[len(b)-1] && c[0] > b[len(b)-1] {
		return 0
	}

	rs := 0
	for i := 0; i < len(b); i++ {
		pp.Println("len(b) = ", len(b))
		pp.Println("i = ", i)
		pp.Println("b = ", b[i])
		aIndex := findIndex(a, b[i], 0, len(a)-1)
		if aIndex > len(a) || aIndex < 0 {
			continue
		}
		cIndex := findIndex(c, b[i], 0, len(c)-1)
		if cIndex > len(c) || cIndex < 0 {
			continue
		}
		pp.Print("aIndex = ", aIndex)
		dupA := a[:aIndex+1]
		pp.Println("dupA = ", dupA)
		pp.Println()
		pp.Print("cIndex = ", cIndex)
		dupC := c[:cIndex+1]
		pp.Println("dupC = ", dupC)
		pp.Println()

		rs += len(dupA) * len(dupC)
	}
	pp.Println(rs)
	return rs
}

func removeDuplicates(el []int) []int {
	e := map[int]bool{}
	redis := []int{}

	for v := range el {
		if e[el[v]] == true {
		} else {
			e[el[v]] = true
			redis = append(redis, el[v])
		}
	}

	return redis
}

func findIndex(ac []int, in int, l, h int) int {
	var m int
	if in > ac[h] {
		return h
	}
	if in < ac[l] {
		return -1
	}

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
	if ac[l] < in || ac[l] > in {
		return l - 1
	}

	return -1
}

func main() {
	filePath := "/Users/namnh/workspace/go/src/github.com/namnhce/my-go-practices/tripple-sum/input.txt"
	output := "/Users/namnh/workspace/go/src/github.com/namnhce/my-go-practices/tripple-sum/output.txt"

	f, err := os.Open(filePath)
	checkError(err)

	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create(output)
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	lenaLenbLenc := strings.Split(readLine(reader), " ")

	lenaTemp, err := strconv.ParseInt(lenaLenbLenc[0], 10, 64)
	checkError(err)
	lena := int(lenaTemp)

	lenbTemp, err := strconv.ParseInt(lenaLenbLenc[1], 10, 64)
	checkError(err)
	lenb := int(lenbTemp)

	lencTemp, err := strconv.ParseInt(lenaLenbLenc[2], 10, 64)
	checkError(err)
	lenc := int(lencTemp)

	arraTemp := strings.Split(readLine(reader), " ")

	var arra []int

	for i := 0; i < int(lena); i++ {
		arraItemTemp, err := strconv.ParseInt(arraTemp[i], 10, 64)
		checkError(err)
		arraItem := int(arraItemTemp)
		arra = append(arra, arraItem)
	}

	arrbTemp := strings.Split(readLine(reader), " ")

	var arrb []int

	for i := 0; i < int(lenb); i++ {
		arrbItemTemp, err := strconv.ParseInt(arrbTemp[i], 10, 64)
		checkError(err)
		arrbItem := int(arrbItemTemp)
		arrb = append(arrb, arrbItem)
	}

	arrcTemp := strings.Split(readLine(reader), " ")

	var arrc []int

	for i := 0; i < int(lenc); i++ {
		arrcItemTemp, err := strconv.ParseInt(arrcTemp[i], 10, 64)
		checkError(err)
		arrcItem := int(arrcItemTemp)
		arrc = append(arrc, arrcItem)
	}

	ans := triplets2(arra, arrb, arrc)

	fmt.Fprintf(writer, "%d\n", ans)

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
