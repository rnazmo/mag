package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const hex = "0123456789ABCDEF"

func init() { rand.Seed(time.Now().UnixNano()) }

// random returns random int in [0,n).
func random(n int) int { return rand.Intn(n) }

func toLower(s []byte) []byte { return bytes.ToLower(s) }
func toUpper(s []byte) []byte { return bytes.ToUpper(s) }

// func genPlane() []byte {
// 	a := make([]byte, 12)
// 	for i := range a {
// 		a[i] = hex[random(len(hex))]
// 	}
// 	return a
// }
// func genHyCo(delimiter byte) []byte {
// 	a := make([]byte, 12+5)
// 	for i := range a {
// 		if i%3 == 2 {
// 			a[i] = delimiter
// 		} else {
// 			a[i] = hex[random(len(hex))]
// 		}
// 	}
// 	return a
// }

func genPlane(p prefix) []byte {
	if p == nil {
		p = make([]byte, 6, 12)
		for i := range p {
			p[i] = hex[random(len(hex))]
		}
	}
	return append(p, []byte{
		hex[random(len(hex))], hex[random(len(hex))],
		hex[random(len(hex))], hex[random(len(hex))],
		hex[random(len(hex))], hex[random(len(hex))],
	}...)
}

func genHyCo(delimiter byte, p prefix) []byte {
	if p == nil {
		// p = make([]byte, 6+2, 12+5) // hituyou?
		p = []byte{
			hex[random(len(hex))], hex[random(len(hex))],
			delimiter,
			hex[random(len(hex))], hex[random(len(hex))],
			delimiter,
			hex[random(len(hex))], hex[random(len(hex))],
		}
	}
	return append(p, []byte{
		delimiter,
		hex[random(len(hex))], hex[random(len(hex))],
		delimiter,
		hex[random(len(hex))], hex[random(len(hex))],
		delimiter,
		hex[random(len(hex))], hex[random(len(hex))],
	}...)
}

// func genDot(delimiter byte) []byte {} // TODO:

func genRandMacAddr(f format, p prefix) []byte {
	// 前提: prefix は validate 済み
	// TODO: format the prefix
	switch f {
	case none:
		// TODO: format prefix to "none"
		return genPlane(p)
	case colon:
		// TODO: format prefix to "colon"
		return genHyCo(':', p)
	case hyphen:
		// TODO: format prefix to "hyphen"
		return genHyCo('-', p)
	}

	log.Fatal("ERROR")
	return nil
}

func main() {
	// TODO: Get OUI-List struct.

	// TODO: Receive config from interactive-stdin or cli option.
	c := newConfig()
	for i := 0; i < c.qty; i++ {
		fmt.Println(string(genRandMacAddr(c.f, c.p)))
	}

	c.p = []byte("0A:0A:0A")
	fmt.Println(string(genRandMacAddr(c.f, c.p)))
	c.f = none
	fmt.Println(string(genRandMacAddr(c.f, c.p)))
}
