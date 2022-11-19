package main

import (
	"fmt"
	"time"
)


func goRoutines()  {
	start := time.Now();
	defer func ()  {
		fmt.Println(time.Since(start))
	}()
	evilNinjas := []string{"Bunty", "Ramesh", "Suresh", "Ravi"}
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja)
	}
	// Main exits even before the goRoutines are finished
	// time.Sleep(time.Second)
}  

func attack(target string)  {
	// might not  execute if not main thread not paused 
	fmt.Println("Attacking ", target)
	time.Sleep(time.Second)
	fmt.Println("Attacked ", target)
}
