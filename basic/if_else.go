package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    // Seed the random number generator for different results each time
    rand.Seed(time.Now().UnixNano())

    if num := rand.Intn(100); num%2 == 0 { // Generate random number, check if even
        fmt.Printf("%d is even\n", num)
    } else {
        fmt.Printf("%d is odd\n", num)
    }
}
   