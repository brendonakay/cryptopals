package main

import (
        "encoding/base64"
        "encoding/hex"
        "fmt"
		"log"
)

func main() {
		const from = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
		decoded, err := hex.DecodeString(from)
		if err != nil {
				log.Fatal(err)
		}
		fmt.Printf("%s\n", decoded)

		data := []byte(string(decoded))
		str := base64.StdEncoding.EncodeToString(data)
		fmt.Println(str)
}
