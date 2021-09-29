// This example demonstrates the difference between unicode.IsNumber and unicode.IsDigit.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	for i := 0; i < 255; i++ {
		l, r := unicode.IsNumber(rune(i)), unicode.IsDigit(rune(i))
		if l != r {
			fmt.Println(l, r, string(rune(i)))
		}

	}
}

//true false ²
//true false ³
//true false ¹
//true false ¼
//true false ½
//true false ¾
