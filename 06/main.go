package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("1.in")

	if err != nil {
		fmt.Println(err)
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	count := 0

	msize := 4
	var marker []rune
	mlen := 0
	found := false

	msgsize := 14
	var msg []rune
	msglen := 0
	found_msg := false
	for fs.Scan() {
		line := string(fs.Text())
		for _, ch := range line {
			count += 1

			marker = append(marker, ch)
			if mlen >= 0 {
				mlen += 1
			}

			msg = append(msg, ch)
			if msglen >= 0 {
				msglen += 1
			}

			if mlen == msize {
				found = true
				for i := 0; i < msize; i++ {
					for x := 0; x < msize; x++ {
						if i == x {
							continue
						}
						if marker[i] == marker[x] {
							found = false
							break
						}
					}
				}
				mlen = msize - 1
				marker = marker[1:msize]
			}
			if found && mlen >= 0 {
				fmt.Printf("Found marker: %d\n", count)
				mlen = -1
			}

			// part 2
			if msglen == msgsize {
				found_msg = true
				for i := 0; i < msgsize; i++ {
					for x := 0; x < msgsize; x++ {
						if i == x {
							continue
						}
						if msg[i] == msg[x] {
							found_msg = false
							break
						}
					}
				}
				msglen = msgsize - 1
				msg = msg[1:msgsize]
			}
			if found_msg && msglen >= 0 {
				fmt.Printf("Found message: %d\n", count)
				msglen = -1
			}
		}

	}

	f.Close()
}
