package main

import (
	"bytes"
	"fmt"
)

func main() {
	//	buf = new bytes.Buffer
	buf := bytes.NewBufferString("aaaa")
	fmt.Println(buf.String())
	buf.WriteString("ccc")
	fmt.Println(buf.String())
	fmt.Println(111, buf.Len())

	testBytes2 := make([]byte, 100)
	buf2 := bytes.NewBuffer(testBytes2)
	fmt.Println(111, buf2.Len())
	fmt.Println(1, buf2.String())
	buf2.Truncate(0)
	buf2.WriteString("heee")
	fmt.Println(2, buf2.String())
	fmt.Println(3, testBytes2)

	testBytesR := make([]byte, 100)
	a, e := buf2.Read(testBytesR)
	fmt.Println(3, a, e)
	fmt.Println(4, buf2.String())
	fmt.Println(5, testBytesR)

	fmt.Println(6, len(testBytesR))
	testBytesR = testBytesR[0:55]
	fmt.Println(6, len(testBytesR))
	fmt.Println(7, cap(testBytesR))

}
