package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Result struct {
	Keysize              int
	NormalHammingDistnce float64
}

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

	// Find smallest average normalized Hamming distance from 4 key sizes
	// TODO - keep in mind KEYSIZE was acquired taking the average of 2 distances.
	var res []Result
	for i := 2; i <= 40; i++ {
		fmt.Println("keysize is:", i)
		a := r[:i]
		fmt.Println("a is:", a)
		b := r[i : 2*i]
		fmt.Println("b is:", b)
		c := r[2*i : 3*i]
		fmt.Println("b is:", c)
		d := r[3*i : 4*i]
		fmt.Println("b is:", d)
		x := float64(hammingDistance(a, b)) / float64(i)
		y := float64(hammingDistance(c, d)) / float64(i)
		res = append(res, Result{
			Keysize:              i,
			NormalHammingDistnce: (x + y) / 2,
		})
	}

	// Wicked awesome sorting closure.
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].NormalHammingDistnce < res[j].NormalHammingDistnce
	})
	for _, s := range res {
		fmt.Println(s.NormalHammingDistnce, s.Keysize)
	}

	// Take the keysize with the smallest average normalized hamming distance
	k := res[0].Keysize
	fmt.Println(k)
}
