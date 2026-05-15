package main

func StackTwo(top []string, bottom []string) []string {
	res := make([]string, 0, len(top)+len(bottom))
	res = append(res, top...)
	res = append(res, bottom...)
	return res
}

func StackAll(blocks [][]string) []string {
	if blocks == nil || len(blocks) == 0 {
		return []string{}
	}

	total := 0
	for _, block := range blocks {
		total += len(block)
	}

	res := make([]string, 0, total)

	for _, block := range blocks {
		res = append(res, block...)
	}

	return res
}