package main

// type octet string
type oui string
type format int
type lettercase int
type prefix []byte

const (
	none   format = iota // xxxxxxxxxxxx
	colon                // xx:xx:xx:xx:xx:xx
	hyphen               // xx-xx-xx-xx-xx-xx
	// TODO: Add support for dot: xxx.xxx.xxx.xxx

	upper lettercase = iota
	lower
)

type config struct {
	q int        // quantity
	l lettercase // lower/upper

	f format // none/colon/hyphen(/dot)

	p prefix

	// TODO: Add support for U/L bit
	// TODO: Add support for I/G bit
	// ul ulbit // https://en.wikipedia.org/wiki/MAC_address#Universal_vs._local_(U/L_bit)
	// ig igbit // https://en.wikipedia.org/wiki/MAC_address#Unicast_vs._multicast_(I/G_bit)
}

func newConfig() config {
	return config{
		q: 3, // TODO: 3 -> 5
		l: upper,
		f: colon,
		p: nil,
		// ul:
	}
}
