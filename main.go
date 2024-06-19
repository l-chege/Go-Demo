package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceTickets = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email    string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}


func main() {
    fmt.Println("Hello World")
}