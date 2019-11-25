package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var Alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Result struct {
	OriginalHex     string
	Rank            int
	ResultByteArray []byte
	RestultString   string
	XorByte         byte
}

// Rudimentary function to determine the "englishness" of a byte array.
// Gives a weight to character bytes and ranks each possible key.
func RankEnglishness(b []byte) int {
	var c int
	if len(b) == 0 {
		c = 0
		return 0
	}
	for i := 0; i < len(b); i++ {
		for j := range Alphabet {
			if b[i] == Alphabet[j] {
				c += 1
			}
		}
	}
	return c
}

func decodeHexString(s string) []Result {
	decoded_s, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DecodeString results: %v\n", decoded_s)
	n := len(decoded_s)

	a := make([]byte, n)
	r := make([]Result, 256)
	for i := 0; i < 256; i++ {
		for j := 0; j < n; j++ {
			a[j] = decoded_s[j] ^ byte(i)
		}
		r = append(r, Result{
			OriginalHex:     s,
			Rank:            RankEnglishness(a),
			ResultByteArray: a[:],
			RestultString:   strconv.QuoteToASCII(string(a)),
			XorByte:         byte(i),
		})
		//fmt.Printf("Key: %s\n", strconv.QuoteToASCII(string(a)))
	}

	return r
}

func main() {

	// Scan file for hex strings
	f, err := os.Open("challenge-data-4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var r []Result
	for s.Scan() {
		r = append(r, decodeHexString(s.Text())...)
	}

	// Wicked awesome sorting closure.
	sort.SliceStable(r, func(i, j int) bool { return r[i].Rank > r[j].Rank })
	for _, s := range r {
		fmt.Println(s.RestultString, s.OriginalHex)
	}
}
