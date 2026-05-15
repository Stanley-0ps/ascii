package main

import "fmt"

const (
	firstASCII = 32
	lastASCII  = 126
	expected   = 95
)

func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return fmt.Errorf("banner is nil")
	}

	if len(banner) != expected {
		return fmt.Errorf("banner has %d entries, expected %d", len(banner), expected)
	}

	for r := rune(firstASCII); r <= lastASCII; r++ {
		art, ok := banner[r]

		if !ok {
			return fmt.Errorf("missing character %q", r)
		}

		if len(art) != 8 {
			return fmt.Errorf("character %q has %d lines, expected 8", r, len(art))
		}
	}

	return nil
}