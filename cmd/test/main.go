package main

import "fmt"

func main() {
	msg := "Greetings\nfrom\nTurkey\n"

	var count int
	for i := 0; i < len(msg); i++ {
		if msg[i] == '\n' {
			count++
		}
	}

	fmt.Println(count)
}
