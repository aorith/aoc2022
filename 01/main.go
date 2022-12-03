package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("1.in")

	if err != nil {
		fmt.Println(err)
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	top := [3]int{0, 0, 0}
	var current int = 0
	var swap int = 0
	for fs.Scan() {
		calories, err := strconv.Atoi(fs.Text())
		if err == nil {
			current += calories
		} else {
			for i := 0; i < len(top); i++ {
				if current > top[i] {
					swap = top[i]
					top[i] = current
					current = swap
				}
			}
			current = 0
		}
	}

	for i := 0; i < len(top); i++ {
		if current > top[i] {
			swap = top[i]
			top[i] = current
			current = swap
		}
	}

	fmt.Println("Top1:", top[0], "Top2:", top[1], "Top3:", top[2], "Total:", top[0]+top[1]+top[2])
	f.Close()
}
