package main

import "strings"

func PadArtRows(rows []string, width int) []string {
	res := make([]string, len(rows))

	for i, row := range rows {
		// Width <= 0 means: return rows unchanged,
		// but still return a NEW slice.
		if width <= 0 {
			res[i] = row
			continue
		}

		if len(row) >= width {
			res[i] = row
			continue
		}

		res[i] = row + strings.Repeat(" ", width-len(row))
	}

	return res
}