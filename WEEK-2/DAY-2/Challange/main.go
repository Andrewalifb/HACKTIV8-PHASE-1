package main

import (
	"fmt"
	"time"
)

func processImage(imageURL string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Image processing completed : %s\n", imageURL)
}
func main() {
	imgUrl := []string{"https://example.com/image1.jpg","https://example.com/image2.jpg","https://example.com/image3.jpg"}
	fmt.Println("Image processiong started, main application continue...")

	for _, link := range imgUrl {
		fmt.Printf("Processing image : %s\n", link)
	}

	for _, link := range imgUrl {
		go processImage(link)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("All image processing completed.")
}