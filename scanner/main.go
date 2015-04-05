package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	var s scanner.Scanner

	str := `
a b "中" 中
c "d"
e f
var 
1.2
2
'ab'
//abc
`
	r := strings.NewReader(str)
	s.Init(r)
	for {
		tok := s.Next()
		if tok == scanner.EOF {
			break
		}
		fmt.Println(s.TokenText() + ":" + scanner.TokenString(tok))
	}

	fmt.Println("----------------")
	r = strings.NewReader(str)
	s.Init(r)
	for {
		tok := s.Scan()
		if tok == scanner.EOF {
			break
		}
		fmt.Println(s.TokenText() + ":" + scanner.TokenString(tok))
	}

}
