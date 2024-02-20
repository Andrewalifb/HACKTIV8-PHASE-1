package main

import (
	// "bufio"
	"fmt"
	"sort"
	"strings"
	// "math/rand"
	// "os"
	// "strings"
)
func penjumlahan(data []float64) float64 {
	var total float64
	for _, value := range data {
			total += value
	}
	return total
}


func rataRata(data []float64) float64 {
	return penjumlahan(data) / float64(len(data))
}


func median(data []float64) float64 {
	copy := make([]float64, len(data))
	copy = append(copy, data...)
	sort.Float64s(copy)
	mid := len(copy) / 2
	if len(copy)%2 == 1 {
			return copy[mid]
	}
	return (copy[mid-1] + copy[mid]) / 2
}

func palindrome(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	for i := 0; i < n/2; i++ {
			if s[i] != s[n-i-1] {
					return false
			}
	}
	return true
}

func hitungKarakter(data []string) bool  {
	x := 0
	o := 0
	var result bool
	for _, val := range data {
		if val == "x" {
			x += 1
		} else if val == "o" {
			o += 1
		}
	}
	if x == o {
		result = true
	} else if x != o {
		result = false
	}
	return result
}

func sortNumbers(numbers []int) {
	swapped := true
	for swapped {
			swapped = false
			for i := 0; i < len(numbers)-1; i++ {
					if numbers[i] > numbers[i+1] {	
							numbers[i], numbers[i+1] = numbers[i+1], numbers[i]			
							swapped = true
					}
			}
	}
}

func perkalian(angka1, angka2 int) int {
	return angka1 * angka2
}
func main() {
	/*NG - CHALLANGE - 2 : Conditional 1*/
// 	var getName string

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Msukan nama : ")

// 	getName, _ = reader.ReadString('\n')

// 	getName = strings.TrimSpace(getName)

// 	randomNumber := rand.Intn(100) + 1

// 	if randomNumber > 80 {
// 		fmt.Println("Selamat", getName, ", anda sangat beruntung jika mendapatkan angka", randomNumber)
// 	} else if randomNumber > 60 && randomNumber <= 80 {
// 		fmt.Println("Selamat", getName, ", anda beruntung jika mendapatkan angka", randomNumber)
// 	} else if randomNumber > 40 && randomNumber <= 60 {
// 		fmt.Println("Mohon Maaf", getName, ", anda kurang beruntung jika mendapatkan angka", randomNumber)
// 	} else if randomNumber <= 40 {
// 		fmt.Println("Mohon Maaf", getName, ", anda sial jika mendapatkan angka", randomNumber)
// 	}

// for i := 1; i <= 15; i ++ {
//  if i % 3 == 0 && i % 5 == 0 {
// 	fmt.Println("Hello World")
// }else if i % 3 == 0 {
// 	fmt.Println("Hello")
// } else if i % 5 == 0 {
// 	fmt.Println("World")
// } else {
// 	fmt.Println(i)
// }
// 	}

/*NG Challenge : Looping 1 */
breakingBad := []map[string]string {
	{"name": "Hank", "Age" : "50", "Job" : "Polisi"},
	{"name": "Heisenberg", "Age" : "52", "Job" : "Ilmuwan"},
	{"name": "Skyler", "Age" : "48", "Job" : "Akuntan"},
}

for i, _ := range breakingBad {
	fmt.Printf("Hai Perkenalkan, Nama saya %s, umur saya %s, dan saya bekerja sebagai %s\n", breakingBad[i]["name"], breakingBad[i]["Age"], breakingBad[i]["Job"])
}


/* NG Challenge : Looping 2 */
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}
	fmt.Println("penjumlahan:", penjumlahan(slice1))
	fmt.Println("Average:", rataRata(slice1))
	fmt.Println("Median:", median(slice1))
	fmt.Println("penjumlahan:", penjumlahan(slice2))
	fmt.Println("Average:", rataRata(slice2))
	fmt.Println("Median:", median(slice2))

/* NG Challenge : Logic 1 - Palindrome*/
	testPalindrome := "katak"
	fmt.Println(palindrome(testPalindrome)) 

/* NG Challenge : Logic 2 - XOXO*/	
xoxo := []string{"x", "o", "x", "o", "x"}
fmt.Println(hitungKarakter(xoxo))

/* NG Challenge : Logic 3 - sort*/	
numbers := []int{5, 3, 8, 2, 1, 4, 7, 6}
sortNumbers(numbers)

/* NG Challenge : Logic 3 - Asterik Level 1*/	
rows1 := 5

for i := 0; i <= rows1; i++ {
	fmt.Println("*")
}

/* NG Challenge : Logic 3 - Asterik Level 2*/	
rows2 := 5

for i := 0; i <= rows2; i++ {
	for j := 0; j <=rows2; j++ {
		fmt.Printf("*")
	}
	fmt.Println()
}

/* NG Challenge : Logic 3 - Asterik Level 3*/	
rows3 := 5
data := ""
for i := 0; i <= rows3; i++ {
		data += "*"
		fmt.Println(data)
}

/* NG Challenge : Logic 3 - Asterik Level 4*/	
rows4 := 5

for i := rows4; i > 0; i-- {
	for j := 1; j <= i; j++ {
			fmt.Print("*")
	}
	fmt.Println()
}

// angka1 := 5
// angka2 := 2

// fmt.Println(perkalian(angka1,angka2))
// func perkalian(angka1, angka2 int) int {
// 	return angka1 * angka2
// }




}




