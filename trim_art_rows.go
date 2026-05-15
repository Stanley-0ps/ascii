package main

import "strings"

func TrimArtRows(rows []string) []string {
	res := make([]string, len(rows))

	for i, row := range rows {
		res[i] = strings.TrimRight(row, " ")
	}

	return res
}