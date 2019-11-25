package main

import (
	"encoding/hex"
	"fmt"
)

func repeatingKeyXorString(k, s string) string {
	var j int
	n := len(s)
	m := len(k)
	a := make([]byte, n)
	b := []byte(k)
	d := []byte(s)
	for i := 0; i < n; i++ {
		a[i] = d[i] ^ b[j%m]
		j++
	}
	return hex.EncodeToString(a)
}
func main() {
	fmt.Println("vim-go")

	m := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	fmt.Println(m)
	fmt.Println(repeatingKeyXorString("ICE", m))
}
