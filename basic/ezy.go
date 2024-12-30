package main

import "fmt"

func main() {
	messagesFromDoris := []string{
		"Yellow is vellow",
		"Blue is blue",
		"Plz help me",
	}
	numMessages := float64(len(messagesFromDoris))

	// Define the cost per message (you need to set a value here)
	const constPerMessage = 1.25 // Example cost per message

	totalCost := constPerMessage * numMessages

	fmt.Printf("Doris spent %.2f\n", totalCost)
}
