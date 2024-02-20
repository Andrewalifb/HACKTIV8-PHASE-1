package main

import (
	"fmt"
	"strings"
)

func rubahCharacter(char rune) rune {
	switch char {
		case 'a':
			return '4'
		case 'e':
			return '3'
		case 'i':
			return '!'
		case 'l':
			return '1'
		case 'n':
			return 'N'
		case 's':
			return '5'
		case 'x':
			return '*'
		default:
			return char 
	}
}

func alayGen(words ...string) {
	sliceToString := strings.Join(words, " ")
	modifiedString := strings.Map(rubahCharacter, sliceToString)
	modifiedWords := strings.Split(modifiedString, strings.Join(words, ""))
	fmt.Println(modifiedWords)
}

func fibonacci(n int) int {
	fibo := make([]int, n+1)
	fibo[0] = 0
	fibo[1] = 1

	for i := 2; i <= n; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
	}
	return fibo[n]
}

func main() {
/* NGC - 4 Function 2 */
alayGen("Hello", "world", "zzz", "ayam", "xiapi", "landasan")
/* NGC - 4 Function 3 */
fmt.Println(fibonacci(6))

}