package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Struct is like a custom class in Java, contains multiple variables or different datatype
type Order struct {
	orderId int
}

func main() {
	// OrderQueue is a channel that is being used as a queue data structure, since it follows FIFO
	orderQueue := make(chan Order) // make created things, and the channel is gonna store data of type Order, meaans if order has 3 fields, then queue's data is gonna have 3 fields.

	// This is a function that runs independently on a thread, since its extended by `go` keyword
	go receptionist(orderQueue)

	// Same as above, a go function that's gonna run independently on a thread
	go chef(orderQueue)

	chefToAnnouncer := make(chan Order)
	go announcer(chefToAnnouncer)

	time.Sleep(1 * time.Second)
}

// The sole purpose of this funtion is to recieve order and store it in the queue (i.e. the channel `orderQueue` in our case)
func receptionist(queue chan Order) {
	for {
		// Generating a random number that'll act as the order number
		randomNum := rand.IntN(100)

		// Adding that number to the queue, so chef starts preparing
		queue <- Order{orderId: randomNum}

		// Still not sure why, but its imp else things don't work
		time.Sleep(100 * time.Microsecond)
	}
}

// The chef funtion that takes exactly 5 secs to prepare any order (khatarnaak chef he is!!)
func chef(queue chan Order) {
	for {
		currentOrder := <-queue
		fmt.Printf("Preparing Order %d", currentOrder.orderId)
		time.Sleep(3 * time.Second)
	}
}

func announcer(queue chan Order) {

}
