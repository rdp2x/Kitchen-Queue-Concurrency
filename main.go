package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Order struct {
	orderId int
}

func main() {
	orderQueue := make(chan Order)

	go receptionist(orderQueue)
	go chef(orderQueue)

	chefToAnnouncer := make(chan Order)
	go announcer(chefToAnnouncer)

	time.Sleep(1 * time.Second)
}

func receptionist(queue chan Order) {
	for {
		randomNum := rand.IntN(100)
		queue <- Order{orderId: randomNum}
		time.Sleep(100 * time.Microsecond)
	}
}

func chef(queue chan Order) {
	for {
		currentOrder := <-queue
		fmt.Printf("Preparing Order %d", currentOrder.orderId)
		time.Sleep(3 * time.Second)
	}
}

func announcer(queue chan Order) {

}
