package main

import (
	"bufio"
	"fmt"
	"os"
)

// go run app.go

func main() {
	fmt.Print("Hello World")
	fmt.Print("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
