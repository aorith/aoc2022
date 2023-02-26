package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var p1 = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var p1r = map[string]int{
	"AX": 3,
	"AY": 6,
	"AZ": 0,
	"BX": 0,
	"BY": 3,
	"BZ": 6,
	"CX": 6,
	"CY": 0,
	"CZ": 3,
}

var p2 = map[string]string{
	// X = lose - X:rock
	// Y = draw - Y:paper
	// Z = win  - Z:scissors
	"AX": "Z",
	"AY": "X",
	"AZ": "Y",
	"BX": "X",
	"BY": "Y",
	"BZ": "Z",
	"CX": "Y",
	"CY": "Z",
	"CZ": "X",
}

func main() {
	f, err := os.Open("1.in")

	if err != nil {
		fmt.Println(err)
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	score := 0
	var sum int = 0
	score2 := 0
	var sum2 int = 0
	var p2choice string
	for fs.Scan() {
		words := strings.Fields(fs.Text())
		fmt.Println(words[0], ":", words[1])
		score = p1[words[1]] + p1r[words[0]+words[1]]
		sum += score
		fmt.Println("Score p1:", score)

		p2choice = p2[words[0]+words[1]]
		score2 = p1[p2choice] + p1r[words[0]+p2choice]
		sum2 += score2
		fmt.Println("Score p2:", score)
	}

	fmt.Println("Total p1:", sum)
	fmt.Println("Total p2:", sum2)
	f.Close()
}
