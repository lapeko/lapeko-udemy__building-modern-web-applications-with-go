package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intChannel := make(chan int)
	defer close(intChannel)

	go getRandFromChannel(intChannel)
	randNumber := <-intChannel
	fmt.Println(randNumber)
}

func getRandFromChannel(ch chan int) {
	ch <- getRandomNumber(100)
}

func getRandomNumber(size int) int {
	return rand.Intn(size)
}
