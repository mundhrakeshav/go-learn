package main

import (
	"fmt"
	"time"
)


func goChannels()  {
	start := time.Now();
	defer func ()  {
		fmt.Println(time.Since(start))
	}()
	smokeSignal := make(chan bool)
	evilNinjas := []string{"Bunty", "Ramesh", "Suresh", "Ravi"}

	// for _, evilNinja := range evilNinjas {
		go attackChannels(evilNinjas[0], smokeSignal)
	// }
	
	// smokeSignal <- true // DeadLock
	
	fmt.Println(<-smokeSignal)
}   

func attackChannels(target string, _signal chan bool)  {
	fmt.Println("Attacking ", target)
	time.Sleep(time.Second)
	fmt.Println("Attacked ", target)
	_signal <- true
}
