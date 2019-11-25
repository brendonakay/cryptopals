package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"sort"
	"strconv"
)

var Alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Result struct {
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

func main() {

	// The encrypted string
	b := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decoded_b, err := hex.DecodeString(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DecodeString results: %v\n", decoded_b)
	n := len(decoded_b)

	a := make([]byte, n)
	r := make([]Result, 256)
	for i := 0; i < 256; i++ {
		for j := 0; j < n; j++ {
			a[j] = decoded_b[j] ^ byte(i)
		}
		r = append(r, Result{
			Rank:            RankEnglishness(a),
			ResultByteArray: a[:],
			RestultString:   strconv.QuoteToASCII(string(a)),
			XorByte:         byte(i),
		})
		//fmt.Printf("Key: %s\n", strconv.QuoteToASCII(string(a)))
	}
	// Wicked awesome sorting closure.
	sort.SliceStable(r, func(i, j int) bool { return r[i].Rank > r[j].Rank })
	for _, s := range r {
		fmt.Println(s.RestultString)
	}
}
