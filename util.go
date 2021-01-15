package main

import (
	"bytes"
	"math/rand"
	"time"
	"unicode"
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
