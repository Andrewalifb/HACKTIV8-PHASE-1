package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Factorial(ch <- chan int) chan int {
	result := make(chan int)
	go func() {
		for n := range ch {
			result <- (n*n)
		}
		close(result) 
	}()
	return result
}

func main() {
	var wg sync.WaitGroup
	randomNumbers := make(chan int)

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	randomInt := r.Intn(100) + 1

	wg.Add(2) 

	go func() {
		defer wg.Done() 

		for i := 1; i <= randomInt; i++ {
			randomNumbers <- i
		}
		close(randomNumbers)
	}()

	go func() {
		defer wg.Done()

		for n := range Factorial(randomNumbers) {
			fmt.Println("Hasil Factorial :",n)
		}
	}()

	wg.Wait()
	fmt.Println("Done")

}

