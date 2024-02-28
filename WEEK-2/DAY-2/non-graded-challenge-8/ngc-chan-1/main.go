package main

import (
	"fmt"
	"strconv"
	"sync"
)

func chanSender(collection chan []string, wg *sync.WaitGroup, totalNumber, totalEvenNumber, totalOddNumber chan int) {
	defer wg.Done()
	var collToSend []string 
	var totalNum, totalEvenNum, totalOddNum int
	for i := 0; i < 100; i++ {
		data := strconv.Itoa(i)
		if i % 3 == 0 && i % 5 == 0 {
			data = data + "FizzBuzz"
			collToSend = append(collToSend, data)
		} else if i % 3 == 0 {
			data = data + "Fizz"
			collToSend = append(collToSend, data)
		} else if i % 5 == 0 {
			data = data + "Buzz"
			collToSend = append(collToSend, data)
		} else {
			collToSend = append(collToSend, data)
		}
		fmt.Println("Send data to channel :", i)

		if i %2 == 0 {
			totalEvenNum += i
		} else if i % 2 != 0 {
			totalOddNum += i
		}
		totalNum += i
	}
	collection <- collToSend
	totalEvenNumber <- totalEvenNum
	totalOddNumber <- totalOddNum
	totalNumber <- totalNum
}

func showData(dataColl string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Print collected data from channel :", dataColl)
}
func main() {
	var wg sync.WaitGroup
	collection := make(chan []string)
	totalNumber := make(chan int)
	totalEvenNumber := make(chan int)
	totalOddNumber := make(chan int)
	wg.Add(1)
	go chanSender(collection, &wg, totalNumber, totalEvenNumber, totalOddNumber)
	getCollection := <- collection
	getEvenNumber := <- totalEvenNumber
	getOddNumber := <- totalOddNumber
	getTotalNumber := <- totalNumber
	for i := 0; i < len(getCollection); i++ {
		wg.Add(1)
		go showData(getCollection[i], &wg)
	}

	wg.Wait()

	fmt.Println("Total penjumlahan semua data integer Even number :", getEvenNumber)
	fmt.Println("Total penjumlahan semua data integer Odd channel :", getOddNumber)
	fmt.Println("Total penjumlahan semua data integer didalam channel :", getTotalNumber)
	fmt.Println("All data processing completed.")
	
}

