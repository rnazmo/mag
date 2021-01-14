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

func (l lettercase) String() string {
	switch l {
	case upper:
		return "UPPERCASE"
	case lower:
		return "lowercase"
	default:
		return "" // TODO: Return error instead of ""(empty string)?
	}
}

func (f format) String() string {
	switch f {
	case none:
		return "None (001122AABBCC)"
	case colon:
		return "Colon (00:11:22:AA:BB:CC)"
	case hyphen:
		return "Hyphen (00-11-22-AA-BB-CC)"
	default:
		return "" // TODO: Return error instead of ""(empty string)?
	}
}

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
