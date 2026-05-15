package main

import (
	"strings"
	"testing"
)

func TestPadArtRows_PadsShortRows(t *testing.T) {
	input := []string{"hi", "there", "a", "b", "c", "d", "e", "f"}
	result := PadArtRows(input, 8)

	for i, row := range result {
		if len(row) != 8 {
			t.Errorf("row %d expected length 8, got %d (%q)", i, len(row), row)
		}
	}
}

func TestPadArtRows_PaddingIsSpaces(t *testing.T) {
	input := []string{"ab", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 5)

	if result[0] != "ab   " {
		t.Errorf("expected %q got %q", "ab   ", result[0])
	}

	if strings.TrimLeft(result[0][2:], " ") != "" {
		t.Errorf("padding must be spaces only")
	}
}

func TestPadArtRows_DoesNotTruncate(t *testing.T) {
	input := []string{"hello world", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 5)

	if result[0] != "hello world" {
		t.Errorf("should not truncate long rows")
	}
}

func TestPadArtRows_ExactWidthUnchanged(t *testing.T) {
	input := []string{"abcd", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 4)

	if result[0] != "abcd" {
		t.Errorf("exact width should remain unchanged")
	}
}

func TestPadArtRows_EmptyRowPadded(t *testing.T) {
	input := []string{"", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 4)

	for i, row := range result {
		if row != "    " {
			t.Errorf("row %d expected 4 spaces, got %q", i, row)
		}
	}
}

func TestPadArtRows_LeadingSpacesPreserved(t *testing.T) {
	input := []string{"  _  ", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 8)

	if !strings.HasPrefix(result[0], "  _  ") {
		t.Errorf("leading spaces must be preserved")
	}
}

func TestPadArtRows_LengthAlwaysPreserved(t *testing.T) {
	input := []string{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh"}
	result := PadArtRows(input, 6)

	if len(result) != len(input) {
		t.Errorf("row count must not change")
	}
}

func TestPadArtRows_DoesNotModifyInput(t *testing.T) {
	input := []string{"hi", "a", "b", "c", "d", "e", "f", "g"}
	copyInput := append([]string{}, input...)

	PadArtRows(input, 10)

	for i := range input {
		if input[i] != copyInput[i] {
			t.Errorf("input must not be modified")
		}
	}
}

func TestPadArtRows_ZeroWidthDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic on width=0: %v", r)
		}
	}()

	input := []string{"hi", "", "", "", "", "", "", ""}
	result := PadArtRows(input, 0)

	if len(result) != 8 {
		t.Errorf("expected same number of rows")
	}
}

func TestPadArtRows_NegativeWidthDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic on negative width: %v", r)
		}
	}()

	input := []string{"hi", "", "", "", "", "", "", ""}
	PadArtRows(input, -5)
}

func TestPadArtRows_AllRowsSameWidth(t *testing.T) {
	input := []string{"a", "bb", "ccc", "d", "ee", "f", "ggg", "hh"}
	result := PadArtRows(input, 10)

	for i, row := range result {
		if len(row) != 10 {
			t.Errorf("row %d expected width 10 got %d", i, len(row))
		}
	}
}