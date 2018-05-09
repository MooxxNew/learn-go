package main

import (
	"fmt"
	"strings"
)

type formatter func(s string) string

func format(s string, fmtrs ...formatter) string {
	for _, fmtr := range fmtrs {
		s = fmtr(s)
	}

	return s
}

func trim(s string) string {
	return strings.Trim(s, " ")
}

func last(s string) string {
	return s[strings.LastIndexByte(s, ' ')+1 : len(s)]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(format(" alan turing ", trim, last, strings.ToUpper))
}
