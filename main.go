package main

import (
	"fmt"
	"log"
)

const hex = "0123456789ABCDEF"

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
	if p == nil || len(p) == 0 {
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
	// TODO: Parse CLI-Option
	//       & Add support for non-interactive mode?

	// TODO: Get OUI-List struct.
	//

	c := newConfig()
	fmt.Println(c)

	if err := c.receiveConfigsInteractively(); err != nil {
		log.Fatal("Failed to receive config interactively", err)
	}
	fmt.Println(c)

	for i := 0; i < c.q; i++ {
		fmt.Println(string(genRandMacAddr(c.f, c.p)))
	}

	c.p = []byte("0A:0A:0A")
	fmt.Println(string(genRandMacAddr(c.f, c.p)))
	c.f = none
	fmt.Println(string(genRandMacAddr(c.f, c.p)))
}
