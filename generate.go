package main

import (
	"bytes"
	"log"
)

const hex = "0123456789ABCDEF"

func genPlane(p prefix) []byte {
	if p == nil || len(p) == 0 {
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

func genRandMacAddr(c config) []byte {
	// 前提: prefix は validate 済み
	// TODO: format the prefix
	var result []byte

	switch c.f {
	case none:
		// TODO: format prefix to "none"
		result = genPlane(c.p)
	case colon:
		// TODO: format prefix to "colon"
		result = genHyCo(':', c.p)
	case hyphen:
		// TODO: format prefix to "hyphen"
		result = genHyCo('-', c.p)
	default:
		log.Fatal("ERROR")
	}

	if c.l == lower {
		result = bytes.ToLower(result)
	}
	return result
}
