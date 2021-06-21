package main

import (
	"fmt"
	"log"
)

func main() {
	// TODO: Parse CLI-Option
	//       & Add support for non-interactive mode?

	// TODO: Get OUI-List struct.
	//

	c := newConfig()
	// fmt.Println(c) // for debug

	if err := c.receiveConfigsInteractively(); err != nil {
		log.Fatal("Failed to receive config interactively", err)
	}
	// fmt.Println(c) // for debug

	for i := 0; i < c.q; i++ {
		fmt.Println(string(genRandMacAddr(c)))
	}
}
