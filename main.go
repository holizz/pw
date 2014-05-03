package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(generate(40))
}

func generate(count int) string {
	s := ""
	for i := 0; i < count; i++ {
		c, err := rand.Int(rand.Reader, big.NewInt(26+26+10))
		if err != nil {
			panic(err)
		}
		d := c.Int64()
		if d < 26 {
			s += string('a' + d)
		} else if d < 26+26 {
			s += string('A' + d - 26)
		} else if d < 26+26+10 {
			s += string('0' + d - (26 + 26))
		}
	}
	return s
}
