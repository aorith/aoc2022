package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func value(str string) int {
	char := []rune(str)
	val := int(char[0])
	if val > 96 && val < 123 {
		return val - 96
	}
	if val > 64 && val < 91 {
		return val - 38
	}
	return 0
}

func main() {
	f, err := os.Open("1.in")

	if err != nil {
		fmt.Println(err)
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	var comp1 []string
	var comp2 []string
	var result []string
	var sum int

	var counter int = 0
	var uniques []string
	p2map := make(map[string]int)
	var p2result []string
	for fs.Scan() {
		counter += 1
		words := strings.Split(fs.Text(), "")
		comp1 = words[0 : len(words)/2]
		comp2 = words[len(words)/2:]
		for i := 0; i < len(comp1); i += 1 {
			if contains(comp2, comp1[i]) {
				if !contains(result, comp1[i]) {
					result = append(result, comp1[i])
				}
			}
		}
		for i := 0; i < len(result); i += 1 {
			sum += value(result[i])
		}
		result = nil

		// p2
		for i := 0; i < len(words); i += 1 {
			if !contains(uniques, words[i]) {
				p2map[words[i]] += 1
				uniques = append(uniques, words[i])
			}
		}
		uniques = nil
		if counter == 3 {
			counter = 0

			for k, e := range p2map {
				if e == 3 {
					p2result = append(p2result, k)
				}
			}
			p2map = make(map[string]int)
		}
	}
	for i := 0; i < len(result); i += 1 {
		sum += value(result[i])
	}
	fmt.Println("Result p1:", sum)

	var sum2 int = 0
	for i := 0; i < len(p2result); i += 1 {
		sum2 += value(p2result[i])
	}
	fmt.Println("Result p2:", sum2)

	f.Close()
}
