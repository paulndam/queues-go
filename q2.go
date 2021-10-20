package main

import (
	"fmt"
	"math/rand"
	"time"
)

// a synchronized queue is a type of queue that needs to be processed in a particular sequence. examples are passengers queue and ticket queue.

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
   )
   //Queue class
   type Queue struct {
	waitPass int
	waitTicket int
	playPass bool
	playTicket bool
	queuePass chan int
	queueTicket chan int
	message chan int
   }

// initializes a queue.
func (queue *Queue) New() {
	queue.message = make(chan int)
	queue.queuePass= make(chan int)
	queue.queueTicket= make(chan int)

	go func() {
		var message int
		for {
		select {
		case message = <-queue.message:
		switch message {
		case messagePassStart:
		queue.waitPass++
		case messagePassEnd:
		queue.playPass = false
		case messageTicketStart:
		queue.waitTicket++
		case messageTicketEnd:
		queue.playTicket = false
		}
		if queue.waitPass > 0 && queue.waitTicket > 0 && !queue.playPass && !queue.playTicket {
		queue.playPass = true
		queue.playTicket = true
		queue.waitTicket--
		queue.waitPass--
		queue.queuePass <- 1
		queue.queueTicket <- 1
		}
		}
		}
		}()
}
	
	   

		// StartTicketIssue starts the ticket issue
func (Queue *Queue) StartTicketIssue() {
	Queue.message <- messageTicketStart
	<-Queue.queueTicket
   }

   // EndTicketIssue ends the ticket issue
func (Queue *Queue) EndTicketIssue() {
	Queue.message <- messageTicketEnd
   }

   //ticketIssue starts and ends the ticket issue
func ticketIssue(Queue *Queue) {
	for {
	// Sleep up to 10 seconds.
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
	Queue.StartTicketIssue()
	fmt.Println("Ticket Issue starts")
	// Sleep up to 2 seconds.
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Println("Ticket Issue ends")
	Queue.EndTicketIssue()
	}
   }

   //StartPass ends the Pass Queue
func (Queue *Queue) StartPass() {
    Queue.message <- messagePassStart
    <-Queue.queuePass
}


//EndPass ends the Pass Queue
func (Queue *Queue) EndPass() {
    Queue.message <- messagePassEnd
}

//passenger method starts and ends the pass Queue
func passenger(Queue *Queue) {
	//fmt.Println("starting the passenger Queue")
	for {
	// fmt.Println("starting the processing")
	// Sleep up to 10 seconds.
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
	Queue.StartPass()
	fmt.Println(" Passenger starts")
	// Sleep up to 2 seconds.
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Println( " Passenger ends")
	Queue.EndPass()
	}
   }

   // main method
func main() {
	var Queue *Queue = & Queue{}
	//fmt.Println(Queue)
	Queue.New()
	fmt.Println(Queue)
	var i int
	for i = 0; i < 10; i++ {
	// fmt.Println(i, "passenger in the Queue")
	go passenger(Queue)
	}
	//close(Queue.queuePass)
	var j int
	for j = 0; j < 5; j++ {
	// fmt.Println(i, "ticket issued in the Queue")
	go ticketIssue(Queue)
	}
	select {}
   }