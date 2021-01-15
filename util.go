package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/pkg/errors"
)

func init() { rand.Seed(time.Now().UnixNano()) }

// random returns random int in [0,n).
func random(n int) int { return rand.Intn(n) }

func toLower(s []byte) []byte { return bytes.ToLower(s) }
func toUpper(s []byte) []byte { return bytes.ToUpper(s) }

func isOnlyInteger(s string) bool {
	if len(s) == 0 { // SpecialCase
		return true
	}
	for _, r := range s {
		if !unicode.IsNumber(r) { // TODO: Use unicode.IsDigit() instead of IsNumber()?
			return false
		}
	}
	return true
}

// NOTE: Alphabet letter case does not matter.
func runeIsIntegerOrAlphabet(r rune) bool {
	return unicode.IsDigit(r) || ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z')
}

// NOTE: Alphabet letter case does not matter.
func isOnlyIntegerOrAlphabet(s string) bool {
	if len(s) == 0 { // SpecialCase
		return true
	}
	for _, r := range s {
		if !runeIsIntegerOrAlphabet(r) {
			return false
		}
	}
	return true
}

// isValidOUI returns whether it is valid as an OUI format.
// Note that it's not whether that the OUI actually exists.
func isValidOUI(s string) bool {
	switch len(s) {
	case 0:
		return false // Dare to make this case (s is "") obvious.
	case 6:
		return isOnlyIntegerOrAlphabet(s) // like "012ABC"
	case 8:
		// first, must be like "01?2A?BC"
		if !(isOnlyIntegerOrAlphabet(s[0:2]) && isOnlyIntegerOrAlphabet(s[3:5]) && isOnlyIntegerOrAlphabet(s[6:8])) {
			return false
		}
		switch {
		case s[2] == ':' && s[5] == ':':
			return true // like "01:2A:BC"
		case s[2] == '-' && s[5] == '-':
			return true // like "01-2A-BC"
		}
	}
	return false
}

func identifyFormatFromOUI(o oui) (format, error) {
	if !isValidOUI(string(o)) {
		return 0, errors.New("wrong format")
	}
	switch {
	case len(o) == 6:
		return none, nil
	case len(o) == 8 && o[2] == ':':
		return colon, nil
	case len(o) == 8 && o[2] == '-':
		return hyphen, nil
	default:
		return 0, errors.New("something wrong")
	}
}

// TODO: Better implementation
func formatOUI(o oui, fWant format) (oui, error) {
	if o == "" { // Special case
		return "", nil
	}
	fNow, err := identifyFormatFromOUI(o)
	if err != nil {
		return "", err
	}
	if fNow == fWant {
		return o, nil
	}
	fmt.Println(fWant, fNow)
	switch fWant {
	case none:
		switch fNow {
		case colon, hyphen:
			// e.g. "01:2A:BC" -> "012ABC"
			return o[0:2] + o[3:5] + o[6:8], nil
		default:
			return "", errors.New("something wrong")
		}
	case colon:
		switch fNow {
		case none:
			// e.g. "012ABC" -> "01:2A:BC"
			return o[0:2] + ":" + o[2:4] + ":" + o[4:6], nil
		case hyphen:
			// e.g. "01-2A-BC" -> "01:2A:BC"
			return oui(strings.ReplaceAll(string(o), "-", ":")), nil
		default:
			return "", errors.New("something wrong")
		}
	case hyphen:
		switch fNow {
		case none:
			// e.g. "012ABC" -> "01-2A-BC"
			return o[0:2] + "-" + o[2:4] + "-" + o[4:6], nil
		case colon:
			// e.g. "01-2A-BC" -> "01:2A:BC"
			return oui(strings.ReplaceAll(string(o), ":", "-")), nil
		default:
			return "", errors.New("something wrong")
		}
	default:
		return "", errors.New("something wrong")
	}
}
