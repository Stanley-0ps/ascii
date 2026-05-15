package main

func cloneRows(rows []string) []string {
	if rows == nil {
		return nil
	}

	cp := make([]string, len(rows))
	copy(cp, rows)
	return cp
}

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	res := make(map[rune][]string)

	for r, art := range base {
		res[r] = cloneRows(art)
	}

	for r, art := range priority {
		res[r] = cloneRows(art)
	}

	return res
}