package main

import (
	"fmt"
	"time"
)


func goBufChannels()  {
	start := time.Now();
	defer func ()  {
		fmt.Println(time.Since(start))
	}()
	smokeSignal := make(chan bool, 1)
	evilNinjas := []string{"Bunty", "Ramesh", "Suresh", "Ravi"}

	smokeSignal <- true // Channel can hold this value as it is a buffered channel
	// smokeSignal <- true // Channel can't hold this value as its size is only 1

	go attackChannels(evilNinjas[0], smokeSignal)
	
	// smokeSignal <- true // DeadLock
	
	fmt.Println(<-smokeSignal, "Log in Driver")
	time.Sleep(time.Second * 6)
	fmt.Println(<-smokeSignal, "Log in Driver2")
	}   

func attackChannelBuf(target string, _signal chan bool)  {
	fmt.Println("Attacking ", target)
	time.Sleep(time.Second)
	fmt.Println("Attacked ", target)
	_signal <- false
}
