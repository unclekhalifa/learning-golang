package main

// "go mod tidy" will automatically look through code and download dependencies if available
// "go run ." will run the code
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
