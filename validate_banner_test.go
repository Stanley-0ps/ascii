package main

import (
	"strings"
	"testing"
)

// buildGoodBanner constructs a perfectly valid banner map in memory.
// No file needed — all 95 printable ASCII chars, each with exactly 8 lines.
func buildGoodBanner() map[rune][]string {
	banner := make(map[rune][]string)
	for r := rune(32); r <= 126; r++ {
		art := make([]string, 8)
		for i := range art {
			art[i] = string(r) + "row"
		}
		banner[r] = art
	}
	return banner
}

func TestValidateBanner_GoodMap(t *testing.T) {
	err := ValidateBanner(buildGoodBanner())
	if err != nil {
		t.Errorf("expected nil for valid banner, got: %v", err)
	}
}

func TestValidateBanner_Nil(t *testing.T) {
	err := ValidateBanner(nil)
	if err == nil {
		t.Error("expected error for nil banner, got nil")
	}
}

func TestValidateBanner_WrongEntryCount(t *testing.T) {
	banner := buildGoodBanner()
	delete(banner, 'A')
	delete(banner, 'B')

	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error for banner with wrong entry count")
	}
}

func TestValidateBanner_MissingSpace(t *testing.T) {
	banner := buildGoodBanner()
	delete(banner, ' ')

	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error for missing space character")
	}
}

func TestValidateBanner_MissingTilde(t *testing.T) {
	banner := buildGoodBanner()
	delete(banner, '~')

	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error for missing tilde (~)")
	}
}

func TestValidateBanner_TooFewLines(t *testing.T) {
	banner := buildGoodBanner()
	banner['A'] = []string{"only", "six", "lines", "here", "not", "eight"}

	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error for too few lines")
	}
	if err != nil && !strings.Contains(err.Error(), "A") {
		t.Errorf("error should mention 'A', got: %v", err)
	}
}

func TestValidateBanner_TooManyLines(t *testing.T) {
	banner := buildGoodBanner()
	banner['Z'] = make([]string, 10)

	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error for too many lines")
	}
}

func TestValidateBanner_ZeroLinesForChar(t *testing.T) {
	banner := buildGoodBanner()
	banner['!'] = []string{}

	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error for zero lines")
	}
}

func TestValidateBanner_EmptyMap(t *testing.T) {
	err := ValidateBanner(map[rune][]string{})
	if err == nil {
		t.Error("expected error for empty map")
	}
}

func TestValidateBanner_ExtraCharacter(t *testing.T) {
	banner := buildGoodBanner()
	banner[rune(200)] = make([]string, 8)

	err := ValidateBanner(banner)
	if err == nil {
		t.Error("expected error for invalid character outside ASCII range")
	}
}

func TestValidateBanner_ErrorMessageIsDescriptive(t *testing.T) {
	banner := buildGoodBanner()
	banner['M'] = []string{"only", "three"}

	err := ValidateBanner(banner)
	if err == nil {
		t.Fatal("expected error")
	}
	if len(err.Error()) < 5 {
		t.Errorf("error too short: %q", err.Error())
	}
}