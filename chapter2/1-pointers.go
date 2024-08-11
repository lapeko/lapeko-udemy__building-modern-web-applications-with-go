package main

import "fmt"

func main() {
	myLine := "Bla bla"
	changeLine(&myLine)
	fmt.Println(myLine)
}

func changeLine(line *string) {
	*line = "New value"
}
