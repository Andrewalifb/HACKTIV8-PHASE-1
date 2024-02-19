package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
/*NG Challenge 1 : Variable 1*/

// A.
var myNum int32
myNum = 50
fmt.Println(myNum)
// B.
var myNum2 float32
myNum = 51
fmt.Println(myNum2)
// C.
var myNumStr string
myNumStr = "50"
fmt.Println(myNumStr)

/*NG Challenge 1 : Variable 2*/
var x int32
x = 5
var y int32
x = 10

fmt.Println("Penjumlahan :", x + y)

/* NG Challenge 1 : CLI */
var getName string


reader := bufio.NewReader(os.Stdin)


fmt.Print("Msukan nama : ")


getName, _ = reader.ReadString('\n')


getName = strings.TrimSpace(getName)


fmt.Println("Hello", getName)

/* NG Challenge 1 : Array & Slice 1 */
people := []string{"Walt", "Jesse", "Skyler", "Saul"}
fmt.Println("Panjang Slice People:", len(people))
people = append(people, "Hank")
people = append(people, "Marie")
fmt.Println("Panjang Slice People:", len(people))
fmt.Println(people)

/* NG Challenge 1 : Array & Slice 2 */
breakingBad := map[string]string {
	"Hank" : "M",
	"Heisenberg" : "M",
	"Skyler" : "F",
}
fmt.Println(breakingBad)
for idx, val := range breakingBad {
	if val == "M" {
		breakingBad[idx] = "Mr"
	} else if val == "F" {
		breakingBad[idx] = "Mrs"
	}
}
fmt.Println(breakingBad)
}
