package main

import (
	"fmt"
	"sync"
)

type Shape struct {
	ShapeType string
	Length    int
	Area      float32
}

func sendData(input []Shape, rectangle, circle, triangle chan Shape, wg *sync.WaitGroup ) {
	defer wg.Done()
	for _, item := range input {
		if item.ShapeType == "RECTANGLE" {
			temp := Shape{ShapeType: item.ShapeType, Length: item.Length, Area: item.Area}
			rectangle <- temp
		} else if item.ShapeType == "CIRCLE" {
			temp := Shape{ShapeType: item.ShapeType, Length: item.Length, Area: item.Area}
			circle <- temp
		} else if item.ShapeType == "TRIANGLE" {
			temp := Shape{ShapeType: item.ShapeType, Length: item.Length, Area: item.Area}
			triangle <- temp
		}
	}
}

func rectangle(data Shape) float32{

	area := (float32(data.Length) * float32(data.Length))
	
	return area
}
func circle(data Shape) float32{

	pi := 3.14
	area := float32(pi) * (float32(data.Length) * float32(data.Length))
	
	return area
}

func triangle(data Shape, tinggi float32) float32{

	
	area := 0.5 * (float32(data.Length) * tinggi)
	
	return area
}


func main() {
	var wg sync.WaitGroup
	input := []Shape{
		{"RECTANGLE", 5, 0.0},
		{"CIRCLE", 5, 0.0},
		{"TRIANGLE", 5, 0.0},
	}

	RECTANGLE := make(chan Shape)
	CIRCLE := make(chan Shape)
	TRIANGLE := make(chan Shape)
	wg.Add(1)
	go sendData(input, RECTANGLE, CIRCLE, TRIANGLE, &wg)
	getRectangle := <-RECTANGLE
	getCircle := <-CIRCLE
	getTriangle := <-TRIANGLE
	
	wg.Wait()
	fmt.Println("Luas Rectangle : ",rectangle(getRectangle))
	fmt.Println("Luas Circle : ",circle(getCircle))
	fmt.Println("Luas Triangle : ",triangle(getTriangle, 5.0))


}