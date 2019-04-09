package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Chuoi ..
type Chuoi interface {
	ToString() string
}

// Spliter ..
type Spliter interface {
	Split(s Chuoi) (Chuoi, Chuoi)
}

type intChuoi struct {
	Value int64
}

func (i intChuoi) ToString() string {
	return fmt.Sprintf("%v", i.Value)
}

type bigIntChuoi struct {
	Value string
}

func (i bigIntChuoi) ToString() string {
	return fmt.Sprintf("%v", i.Value)
}

func (i bigIntChuoi) Split() (*bigIntChuoi, *bigIntChuoi) {
	r1, r2 := "", ""
	for idx, c := range i.Value {
		if idx == 0 {
			fr1, fr2 := splitFirstCharacter(string(c))
			r1 += fr1
			r2 += fr2
			continue
		}
		if string(c) == "4" {
			r1 += "2"
			r2 += "2"
		} else {
			r1 += "0"
			r2 += string(c)
		}
	}

	return &bigIntChuoi{Value: r1}, &bigIntChuoi{r2}
}

type bigIntSpliter struct {
}

func (h bigIntSpliter) Split(v Chuoi) (Chuoi, Chuoi) {
	val, ok := v.(bigIntChuoi)

	if !ok {
		return nil, nil
	}

	rs1, rs2 := val.Split()

	return rs1, rs2
}

type intSpliter struct {
}

func (h intSpliter) Split(v Chuoi) (Chuoi, Chuoi) {
	return nil, nil
}

func splitFirstCharacter(c string) (string, string) {
	r1, r2 := "", ""
	if string(c) == "0" {
		return "", ""
	}
	if string(c) == "1" {
		return "1", ""
	}

	if string(c) == "2" {
		return "1", "1"
	}

	if string(c) == "3" {
		return "1", "2"
	}

	if string(c) == "4" {
		return "2", "2"
	}

	if string(c) == "5" {
		return "2", "3"
	}

	if string(c) == "6" {
		return "3", "3"
	}

	if string(c) == "7" {
		return "2", "5"
	}

	if string(c) == "8" {
		return "2", "6"
	}

	if string(c) == "9" {
		return "3", "6"
	}

	return r1, r2
}

func main() {
	filePath := "input.txt"

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReaderSize(f, 1024*1024*1024)
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		panic(err)
	}

	q := int32(qTemp)

	spliter := bigIntSpliter{}

	for i := 0; i < int(q); i++ {
		input := readLine(reader)
		var c Chuoi
		c = bigIntChuoi{Value: (input)}
		rs1, rs2 := spliter.Split(c)
		fmt.Printf("Case #%v: %v %v\n", i+1, rs1.ToString(), rs2.ToString())

	}

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
