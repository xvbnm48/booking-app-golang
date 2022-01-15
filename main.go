package main

import "fmt"

func main() {
	var conferenceName = "GO Conference"
	const conferenTickets = 50
	var remainingTickets = 50
	fmt.Println("welcome to our", conferenceName, "booking application")
	fmt.Println("we have total of ,", conferenTickets, "tickets and ,", remainingTickets, "are stil available")
	fmt.Println("get your tickets here to attend")

	fmt.Println(conferenceName)
}
