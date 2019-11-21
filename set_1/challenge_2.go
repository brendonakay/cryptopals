package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const a = "1c0111001f010100061a024b53535009181c"
	decoded_a, err := hex.DecodeString(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", decoded_a)

	const b = "686974207468652062756c6c277320657965"
	decoded_b, err := hex.DecodeString(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", decoded_b)

	n := len(decoded_a)
	c := make([]byte, n)
	for i := 0; i < n; i++ {
		c[i] = decoded_a[i] ^ decoded_b[i]
	}

	o := hex.EncodeToString(c)

	fmt.Printf("%v\n", o)

}
