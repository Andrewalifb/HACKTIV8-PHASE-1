package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func checkAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	count := make(map[rune]int)

	for _, char := range str1 {
		count[char]++
	}

	for _, char := range str2 {
		count[char]--
		if count[char] < 0 {
			return false
		}
	}

	return true
}

func checkInput(input string) {
	
	if len(input) > 10 {
		defer func() {
			recover() 
			fmt.Println("Input tidak boleh lebih dari 10 huruf")
	
		}()
		panic("error")
	}

	for _, char := range input {
		if !unicode.IsLetter(char) {
			defer func() {
				recover() 
				fmt.Println("Input hanya boleh mengandung huruf")	
	
				}()
			panic("error")
		}
		if unicode.IsSpace(char) {
			defer func() {
				recover() 
				fmt.Println("Input tidak boleh memiliki spasi")	
		
			}()
			panic("error")
		}
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first string: ")
	str1, err := reader.ReadString('\n')
	if err != nil {
		defer func() {
			recover() 
			fmt.Println("Invalid Input pada first string")	
		}()
		panic(err)
	} else {
		checkInput(str1)
	}

	fmt.Print("Enter second string: ")
	str2, err := reader.ReadString('\n')
	if err != nil {
		defer func() {
			recover() 
			fmt.Println("Invalid Input pada first string")	
		}()
		panic(err)
	} else {
		checkInput(str2)
	}

	str1 = strings.TrimSpace(str1)

	str2 = strings.TrimSpace(str2)

	isAnagram := checkAnagram(str1, str2)

	if isAnagram {
		fmt.Println("String adalah anagram.")
	} else {
		fmt.Println("String bukan anagram anagrams.")
	}
}


