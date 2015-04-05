package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	b, _ := br.Peek(5)
	fmt.Printf("%s\n", b)
	// ABCDE
	fmt.Println(1, len(b), cap(b))

	b[0] = 'a'
	b, _ = br.Peek(5)
	fmt.Printf("%s\n", b)
	// aBCDE
}
