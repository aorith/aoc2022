package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	start int
	end   int
}

func (p1 Pair) Contains(p2 Pair) bool {
	if p1.start <= p2.start && p1.end >= p2.end {
		return true
	}
	if p2.start <= p1.start && p2.end >= p1.end {
		return true
	}
	return false
}

func (p1 Pair) Overlaps(p2 Pair) bool {
	for i := p1.start; i <= p1.end; i++ {
		if i >= p2.start && i <= p2.end {
			return true
		}
	}
	for i := p2.start; i <= p2.end; i++ {
		if i >= p1.start && i <= p1.end {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("1.in")

	if err != nil {
		fmt.Println(err)
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	counter := 0
	counter2 := 0
	for fs.Scan() {
		words := strings.Split(fs.Text(), ",")
		raw_pair1 := strings.Split(words[0], "-")
		raw_pair2 := strings.Split(words[1], "-")

		p1_start, err := strconv.Atoi(raw_pair1[0])
		if err != nil {
			panic(err)
		}
		p1_end, err := strconv.Atoi(raw_pair1[1])
		if err != nil {
			panic(err)
		}

		p2_start, err := strconv.Atoi(raw_pair2[0])
		if err != nil {
			panic(err)
		}
		p2_end, err := strconv.Atoi(raw_pair2[1])
		if err != nil {
			panic(err)
		}

		p1 := Pair{
			start: p1_start,
			end:   p1_end,
		}
		p2 := Pair{
			start: p2_start,
			end:   p2_end,
		}

		if p1.Contains(p2) {
			counter += 1
		}

		if p1.Overlaps(p2) {
			counter2 += 1
		}
	}
	fmt.Println("Result p1:", counter)
	fmt.Println("Result p2:", counter2)

	f.Close()
}
