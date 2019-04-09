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

func exercise3(maxPrime int64, ciphertext []int64, lenCipher int64) string {
	p := [][]int64{}
	for i := 0; i < len(ciphertext); i++ {
		n1, n2 := multiplyBy(ciphertext[i])

		if i > 0 && n1 != p[i-1][1] {
			p = append(p, []int64{n2, n1})
		} else {
			p = append(p, []int64{n1, n2})
		}
		if i > 0 && n1 != p[i-1][1] && n2 != p[i-1][1] {
			i = 0
			p = [][]int64{}
			n1, n2 = multiplyBy(ciphertext[0])
			p = append(p, []int64{n2, n1})
		}
	}

	p2 := []int64{}
	for i, s := range p {
		if i > 0 {
			p2 = append(p2, s[1])
		} else {
			p2 = append(p2, s[0])
			p2 = append(p2, s[1])
		}
	}
	rs := ""
	for _, s1 := range p2 {
		index := 0
		for _, s2 := range unique(p2) {
			if s1 > s2 {
				index++
			}
		}
		rs += fmt.Sprintf("%c", index+65)
	}

	return rs
}

func multiplyBy(in int64) (int64, int64) {
	inSqrt := math.Sqrt(float64(in))
	for i := int64(inSqrt); i > 0; i-- {
		if in%i == 0 {
			return i, in / i
		}
	}
	return 0, 0
}

func main() {
	filePath := "input.txt"

	f, err := os.Open(filePath)
	checkError(err)

	reader := bufio.NewReaderSize(f, 1024*1024*1024)
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	q := int32(qTemp)

	for i := 0; i < int(q); i++ {

		st := strings.Split(readLine(reader), " ")
		var ciphertext []int64
		maxPrime, err := strconv.ParseInt(st[0], 10, 64)
		checkError(err)

		lenCipher, err := strconv.ParseInt(st[1], 10, 64)
		checkError(err)

		st2 := strings.Split(readLine(reader), " ")
		for _, s := range st2 {
			c, err := strconv.Atoi(s)
			checkError(err)

			ciphertext = append(ciphertext, int64(c))
		}

		rs := exercise3(maxPrime, ciphertext, lenCipher)

		fmt.Printf("Case #%v: %v\n", i+1, rs)

	}

}

func unique(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	list := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
