package main

import "fmt"

func main() {
	str := "a国\u65e7\U000065e7"
	fmt.Println("len(", str, ")=", len(str)) //一个汉字在UTF-8中占3个字节
	fmt.Println("str[0]=", str[0])           //l
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
	for i, s := range str {
		fmt.Println(i, "Unicode(", s, ") string=", string(s))
	}
	r := []rune(str)
	fmt.Println("rune=", r)
	for i := 0; i < len(r); i++ {
		fmt.Println("r[", i, "]=", r[i], "string=", string(r[i]))
	}
}
