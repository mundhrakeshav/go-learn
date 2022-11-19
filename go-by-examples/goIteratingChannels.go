package main

import (
	"fmt"
	"math/rand"
	"time"
)


func goIterChannels()  {
	
	var numRounds uint8 = 3;
	
	smokeSignal := make(chan string);
	go throwNinjaStar(smokeSignal, numRounds)
	for  {
		message, open := <-smokeSignal
		if !open {
			break
		}
		fmt.Println(message)
	}

	// for message:= range smokeSignal {
	// 	fmt.Println(message)
	// }
}   

func throwNinjaStar(_signal chan string, numRounds uint8)  {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < int(numRounds); i++ {
		score := rand.Intn(10)
		_signal <- fmt.Sprint("Score is:", score)
	}
	close(_signal)
}
