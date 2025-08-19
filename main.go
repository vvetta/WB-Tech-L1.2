package main

import "fmt" 

func main() {
	ints := []int{2, 4, 6, 8, 10}

	results := make(chan int, len(ints))

	for _, i := range ints {
		go func(n int){
			results <- n*n	
		}(i)
	}

	for range ints {
		fmt.Println(<-results)
	}
}
