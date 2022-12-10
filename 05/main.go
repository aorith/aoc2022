package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items []rune
}

func (s *Stack) append(r rune) {
	s.items = append(s.items, r)
}

func (s *Stack) insert(r rune) {
	s.items = append([]rune{r}, s.items...)
}

func (s *Stack) pop() rune {
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func main() {
	f, err := os.Open("1.in")

	if err != nil {
		fmt.Println(err)
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	buf := bytes.Buffer{}
	// find the stacks
	var stack_idx []int
	found := false
	for fs.Scan() {
		line := string(fs.Text())
		fmt.Fprintln(&buf, line) // to parse the input twice...

		if len(line) < 2 || found {
			continue
		}

		for i, b := range line {
			found = true
			c := string(b)
			_, err := strconv.Atoi(c)
			if err != nil && c != " " {
				found = false
				stack_idx = nil
				break
			} else if err == nil {
				stack_idx = append(stack_idx, i)
			}
		}
	}

	// fill the stacks
	stacks := make([]Stack, len(stack_idx))
	stacks2 := make([]Stack, len(stack_idx))
	temp_stack := Stack{}
	s2 := bufio.NewScanner(&buf)
	for s2.Scan() {
		line := string(s2.Text())
		if len(line) == stack_idx[len(stack_idx)-1]+2 {
			for i, n := range stack_idx {
				if string(line[n+1]) == "]" {
					stacks[i].insert(rune(line[n]))
					stacks2[i].insert(rune(line[n]))
				}
			}
		}

		if len(line) > 1 && string(line[0]) == "m" {
			words := strings.Split(line, " ")
			moves, _ := strconv.Atoi(words[1])
			from, _ := strconv.Atoi(words[3])
			to, _ := strconv.Atoi(words[5])
			temp_stack.items = nil
			for i := 0; i < moves; i++ {
				stacks[to-1].append(stacks[from-1].pop())
				temp_stack.append(stacks2[from-1].pop())
			}

			temp_sz := len(temp_stack.items)
			for i := 0; i < temp_sz; i++ {
				stacks2[to-1].append(temp_stack.pop())
			}
		}
	}

	// Print result
	fmt.Print("Part 1: ")
	for i := 0; i < len(stack_idx); i++ {
		fmt.Print(string(stacks[i].items[len(stacks[i].items)-1]))
	}
	fmt.Println()
	fmt.Print("Part 2: ")
	for i := 0; i < len(stack_idx); i++ {
		fmt.Print(string(stacks2[i].items[len(stacks2[i].items)-1]))
	}
	fmt.Println()

	f.Close()
}
