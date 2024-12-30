package main

import "fmt"

func main() {
	names := [3]string{"alice", "bob", "charlie"}
	var age [3]int16 // Note: You likely want to initialize the ages

	// Method 1: Using a for loop with index
	fmt.Println("Using for loop with index:")
	for i := 0; i < len(names); i++ {
		fmt.Printf("Index: %d, Name: %s\n", i, names[i])
	}

	// Method 2: Using a for...range loop (most similar to Python's enumerate)
	fmt.Println("\nUsing for...range loop:")
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	// Example showing how you might associate names and ages:
	age[0] = 30
	age[1] = 25
	age[2] = 35

	fmt.Println("\nNames and Ages (using for...range):")
	for i, name := range names {
        if i < len(age){
            fmt.Printf("Index: %d, Name: %s, Age: %d\n", i, name, age[i])
        }
	}
    fmt.Println("\nNames and Ages (using for loop):")
	for i := 0; i < len(names); i++ {
        if i < len(age){
		    fmt.Printf("Index: %d, Name: %s, Age: %d\n", i, names[i], age[i])
        }
	}
}
