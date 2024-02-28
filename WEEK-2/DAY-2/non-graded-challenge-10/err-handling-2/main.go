package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pilih operasi matematika")
	fmt.Println("Penjumlahan (+) kode : a")
	fmt.Println("Pengurangan (-) kode : b")
	fmt.Println("Perkalian   (*) kode : c")
	fmt.Println("Pembagian   (/) kode : d")
	fmt.Print("Enter Kode: ")
  kode, err := reader.ReadString('\n')
	if err != nil {
		defer func() {
			recover() 
			fmt.Println("Invalid Input pada kode")	
		}()
		panic(err)
  }
	kode = strings.TrimRight(kode, "\r\n")
	if kode == "a" {
		kode = "+"
	} else if kode == "b" {
		kode = "-"
	} else if kode == "c" {
		kode = "*"
	} else if kode == "d" {
		kode = "/"
	} else {
		fmt.Println("Anda salah memasukkan kode")
		return
	}

  fmt.Print("Masukkan angka pertama: ")
  str1, err := reader.ReadString('\n')
  if err != nil {
		defer func() {
			recover() 
			fmt.Println("Invalid Input pada first integer")	
		}()
		panic(err)
  }


  str1 = strings.TrimRight(str1, "\r\n")

  num1, err := strconv.Atoi(str1)
  if err != nil {

		defer func() {
			recover() 
			fmt.Println("Input harus berupa integer")	
		}()
		panic(err)
  }

  fmt.Print("Masukkan angka kedua: ")
  str2, err := reader.ReadString('\n')
  if err != nil {

		defer func() {
			recover() 
			fmt.Println("Invalid Input pada second integer")	
		}()
		panic(err)
  }


  str2 = strings.TrimRight(str2, "\r\n")

  num2, err := strconv.Atoi(str2)
  if err != nil {

		defer func() {
			recover() 
			fmt.Println("Input harus berupa integer")	
		}()
		panic(err)
  }

	if kode == "+" {
		fmt.Printf("Hasil dari %d + %d adalah : %d", num1, num2, (num1 + num2))
	} else if kode == "-" {
		fmt.Printf("Hasil dari %d - %d adalah : %d", num1, num2, (num1 - num2))
	} else if kode == "*" {
		fmt.Printf("Hasil dari %d * %d adalah : %d", num1, num2, (num1 * num2))
	} else if kode == "/" {
		fmt.Printf("Hasil dari %d / %d adalah : %d", num1, num2, (num1 / num2))
	} 

}
