package main

import (
	"strings"
	"testing"
)

func TestTrimArtRows_RemovesTrailingSpaces(t *testing.T) {
	input := []string{"hello   ", "world  ", "foo", "bar   ", "", "  ", "a ", " b "}
	result := TrimArtRows(input)

	want := []string{"hello", "world", "foo", "bar", "", "", "a", " b"}

	for i := range want {
		if result[i] != want[i] {
			t.Errorf("row %d got %q want %q", i, result[i], want[i])
		}
	}
}

func TestTrimArtRows_LengthUnchanged(t *testing.T) {
	input := []string{"a  ", "b", "c   ", "d", "e  ", "f", "g   ", "h"}
	result := TrimArtRows(input)

	if len(result) != len(input) {
		t.Errorf("length changed")
	}
}

func TestTrimArtRows_AllEmptyRows(t *testing.T) {
	input := []string{"", "", "", "", "", "", "", ""}
	result := TrimArtRows(input)

	for i, row := range result {
		if row != "" {
			t.Errorf("row %d should be empty", i)
		}
	}
}

func TestTrimArtRows_AllSpaceRows(t *testing.T) {
	input := []string{"   ", "  ", " ", "    ", "     ", "  ", "   ", " "}
	result := TrimArtRows(input)

	for i, row := range result {
		if row != "" {
			t.Errorf("row %d should become empty string", i)
		}
	}
}

func TestTrimArtRows_NoTrailingSpaces(t *testing.T) {
	input := []string{"_", "| |", "|_|", "", " _", "| |", "|_|", ""}
	result := TrimArtRows(input)

	for i := range input {
		if result[i] != input[i] {
			t.Errorf("row %d changed unexpectedly", i)
		}
	}
}

func TestTrimArtRows_DoesNotModifyInput(t *testing.T) {
	input := []string{"hi   ", "there  ", "a   ", "b  ", "c ", "d  ", "e  ", "f   "}
	copyInput := append([]string{}, input...)

	TrimArtRows(input)

	for i := range input {
		if input[i] != copyInput[i] {
			t.Errorf("input modified")
		}
	}
}

func TestTrimArtRows_ReturnsNewSlice(t *testing.T) {
	input := []string{"a  ", "b  ", "c  ", "d  ", "e  ", "f  ", "g  ", "h  "}
	result := TrimArtRows(input)

	result[0] = "MUTATED"

	if input[0] == "MUTATED" {
		t.Error("must return new slice")
	}
}

func TestTrimArtRows_MidRowSpacesPreserved(t *testing.T) {
	input := []string{"| _ |   ", "| | |  ", "|___|", "", "", "", "", ""}
	result := TrimArtRows(input)

	if !strings.Contains(result[0], "| _ |") {
		t.Errorf("internal spaces lost")
	}
	if !strings.Contains(result[1], "| | |") {
		t.Errorf("internal spaces lost")
	}
}

func TestTrimArtRows_EmptySlice(t *testing.T) {
	result := TrimArtRows([]string{})

	if result == nil {
		t.Error("must not return nil")
	}
	if len(result) != 0 {
		t.Error("must return empty slice")
	}
}