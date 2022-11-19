package main

import (
	"fmt"
	"time"
)


func goChannels2()  {
	start := time.Now();
	defer func ()  {
		fmt.Println(time.Since(start))
	}()
	smokeSignal := make(chan string)
	evilNinjas := []string{"Bunty", "Ramesh", "Suresh", "Ravi"}

	// smokeSignal <- "Value2"

	// for _, evilNinja := range evilNinjas {
		go attackChannel21(evilNinjas[0], smokeSignal)
		// go attackChannel22(evilNinjas[1], smokeSignal)
	// }
	
	// smokeSignal <- true // DeadLock
	time.Sleep(time.Second * 5)
	fmt.Println(<-smokeSignal, "In Driver")
}   

func attackChannel21(target string, _signal chan string)  {
	// fmt.Println(<-_signal, "AttackChannel21")
	fmt.Println("Attacking ", target)
	time.Sleep(time.Second)
	fmt.Println("Attacked ", target)
	_signal <- "true"
}

func attackChannel22(target string, _signal chan string)  {
	fmt.Println("Attacking ", target)
	fmt.Println(<-_signal, "AttackChannel22")
	time.Sleep(time.Second)
	fmt.Println("Attacked ", target)
	// _signal <- true
}
