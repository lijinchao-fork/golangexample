package main

import (
	"fmt"
	"os"
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

	fmt.Println("----------------")
	f, _ := os.OpenFile("main.go", os.O_APPEND, 0666)
	defer f.Close()
	s.Init(f)
	for {
		tok := s.Scan()
		if tok == scanner.EOF {
			break
		}
		fmt.Println("main.go", s.TokenText()+":"+scanner.TokenString(tok))
	}

}
