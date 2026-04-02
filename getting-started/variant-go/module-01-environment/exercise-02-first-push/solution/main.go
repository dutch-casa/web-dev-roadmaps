package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Your Name"
	today := time.Now().Format("January 2, 2006")
	fmt.Printf("Hello, I'm %s. Today is %s.\n", name, today)
}
