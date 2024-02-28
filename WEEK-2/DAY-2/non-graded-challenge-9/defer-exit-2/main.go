package main

import (
	"fmt"
	"sync"
)

type Circle struct {
	Area          float64
	Circumference float64
}

func calcArea(pi float64, radius int) float64 {

	area := (2 * pi) * float64(radius)

	return area
}

func calcCircumference(pi float64, radius int) float64 {

	circumference := pi * float64(radius) * float64(radius)

	return circumference
}

func main() {
	var wg sync.WaitGroup
	result := make(chan Circle)
	radius := []int{5, 6, 7, 8, 9, 10, 11, 12}
	pi := 3.14

	wg.Add(2) 
	go func() {
		defer wg.Done() 

		for i := 0; i < len(radius); i++ {
		
			temp := Circle{
				Area: calcArea(pi, radius[i]),
				Circumference: calcCircumference(pi, radius[i]),
			}
				result <- temp
		}
		close(result)
		
	}()

	go func() {
	defer wg.Done()

		for n := range result {
			fmt.Printf("Area : %.2f\n", n.Area)
			fmt.Printf("Circumference : %.2f\n", n.Circumference)
		}

	}()
	
	wg.Wait()
	fmt.Println("Done")
}