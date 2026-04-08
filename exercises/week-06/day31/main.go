package main

import "fmt"

func summarize(values []any) string {
	var total int
	var strList []string
	var trueCount int

	for _, v := range values {
		switch val := v.(type) {
		case bool:
			if val {
				trueCount++
			}
		case string:
			strList = append(strList, val)
		case int:
			total += val
		}
	}
	return fmt.Sprintf("total %d, words: %v, trues: %d", total, strList, trueCount)
}

func main() {
	r := summarize([]any{1, "hello", true, 4, "world", false, 10, true})
	fmt.Println(r)
}
