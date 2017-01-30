package main
//Shortest solution for the GOlang tour maps exercise

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	chunks := strings.Split(s, " ")
	for _, v := range(chunks) {
		cnt, _ := res[v]
		res[v] = cnt+1
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
