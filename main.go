package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Struct is like a custom class that has multiple objects.
type Order struct {
	orderId int
}

func main() {

	// Basically a queue that holds all the orders
	orderQueue := make(chan Order)

	// The `go` keyword here is used so the function run on a dedicated thread.
	// The receptionist will just make sure the input recieved lands in the order queue.
	go receptionist(orderQueue)

	// This is a channel that'll act a a channel only cz this time it'll be used to pass data from one goroutine(chef) to another goroutine(announcer)
	// This time announcement channel will carry the order that are prepared and send it to the announce. So internally the announcer is gonna work same as chef, value dosent matter, its presence matters, so the monent it sees any data in announcement channel, it'll say/print the value it see.
	announcement := make(chan Order)

	// Same here, the chef will prepare the order once there he ensures there are items in the order queue. Think of this as the chef got to know there's an order so he stands up and starts the work.
	go chef(orderQueue, announcement)

	go announcer(announcement)

	select {}
}

func receptionist(queue chan Order) {
	for {

		// Creats a random number
		randomNum := rand.IntN(100)

		// Adds the number to the order queue. Something like what I see a McDonalds.
		queue <- Order{orderId: randomNum}

		// Without a tiny pause, the computer's CPU runs so fast that the receptionist will try to push millions of orders into the channel a second. While it does technically work without it, adding a small delay just acts as a human-like throttle so your terminal doesn't instantly get flooded with millions of numbers before the chef can even blink.
		time.Sleep(100 * time.Microsecond)
	}
	// `for { ... }` in Go is simply the equivalent of `while (true)` or an endless loop in other languages. Because you want the restaurant to stay open forever until you press Ctrl + C, you just leave the conditions blank. It loops infinitely.

	// Also idk is I have to return anything. Most probably not, cz I'm just changing the order queue, and not doing anything in a new queue/channel
}

func chef(queue chan Order, messageForAnnouncer chan Order) {

	// Again a loop, idk how its working, but its there cz thigns are to be done repeated.
	for {
		// currentOrder will hold the first value from the queue. Basically, the order number isn't going to matter here, its presence matters, cz rn the chef sees a number and starts preparing the food. (new restaurent so just one item available for now!!)
		currentOrder := <-queue

		// Before staring to prepare, he'll say that what order number he's preparing!
		fmt.Printf("Preparing Order %d\n", currentOrder.orderId)

		// Each order takes 3 seconds to get prepared. (Magical chef, wahhhhh)
		time.Sleep(3 * time.Second)

		// This will add the prepared order number to messageForAnnouncer/announcement and return it. but i'm not sure that it needs this returning!!
		messageForAnnouncer <- currentOrder
	}
}

func announcer(queue chan Order) {
	// idk how the announcement channel will be used cz never done it. I've used channels as queues, not as a means of data transfer so idk. So chances are the announcer's code written might be incorrect.
	for {
		//
		completedOrder := <-queue
		fmt.Printf("Order Number %d Ready to Serve!\n", completedOrder.orderId)
		time.Sleep(3 * time.Microsecond)
	}
}
