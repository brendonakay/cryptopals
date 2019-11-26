package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func hammingDistance(a, b string) int {
	// props to https://stackoverflow.com/a/40309527
	if len(a) != len(b) {
		return 0
	}
	c := []byte(a)
	d := []byte(b)
	diff := 0
	for i := 0; i < len(a); i++ {
		b1 := c[i]
		b2 := d[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (b1 & mask) != (b2 & mask) {
				diff++
			}
		}
	}
	return diff
}

// Set 1 Challenge 6 Cryptopals
// TODO
//	- Normalize k key length iteration over cipher text
//	- ...
func main() {
	fmt.Println(hammingDistance("this is a test", "wokka wokka!!!"))

	// Scan file for b64 encoded string
	f, err := os.Open("set_1/challenge_6/challenge-6-data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var r string
	for s.Scan() {
		r += strings.TrimRightFunc(s.Text(), func(c rune) bool {
			// Account for either style line ending
			return c == '\r' || c == '\n'
		})
	}

	fmt.Println(r)

	// Find smallest Hamming distance from key sizes
	keySizes := make([]int, 40)
	for i := 1; i < len(keySizes); i++ {
		a := r[:i]
		b := r[i : i+i]
		keySizes[i] = hammingDistance(a, b)
	}

	fmt.Println(keySizes)
}
