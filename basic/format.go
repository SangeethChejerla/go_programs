package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int = 100          // Example limit
	var costPerSMS float64 = 0.02         // Example cost
	var hasPermission bool = true        // Example permission status
	var username string = "AliceBob123" // Example username

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)
}
