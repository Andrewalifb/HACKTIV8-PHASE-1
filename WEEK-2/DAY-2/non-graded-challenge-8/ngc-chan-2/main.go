package main

import (
	"fmt"
	"sync"
)

func SumSquare(num int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sumNumber, squareNumber int
	fmt.Println("All data processing started.")
	for i := 0; i <= num; i++ {
    sumNumber += i
	}
	squareNumber = sumNumber * sumNumber
	result <- squareNumber
}

func SquareSum(num int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sumNumber int
	for i := 0; i <= num ; i++ {
		sumNumber += i * i
	}
	result <- sumNumber
}

func main() {
	var wg sync.WaitGroup
  // NG CHallenge 8 : Channel 2
  num := 100
  resultSumSquare := make(chan int) 
  resultSquareSum := make(chan int)

	wg.Add(2)
	go SumSquare(num, resultSumSquare, &wg)
	go SquareSum(num, resultSquareSum, &wg)
	getSumSquare := <- resultSumSquare
	getSquareSum := <- resultSquareSum

	wg.Wait()

	fmt.Println("SumSquare result :", getSumSquare) // 3025
	fmt.Println("SquareSum result :", getSquareSum) // 385
	fmt.Println("All data processing completed.")
	
	
}

