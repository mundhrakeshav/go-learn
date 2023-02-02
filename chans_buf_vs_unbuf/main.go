package main

import (
	"time"
)

type T struct {
	a byte;
	b bool;
}

func send(i int, count *int, ch chan <- *T)  {
	t := &T{a: byte(i)}
	*count = *count+1;
	ch <- t
	t.b = true;
}

func main() {
	count := 0;
	arr := make([]T, 5)
	// If this channel isn't buffered the `t.b = true;` wont take place as 
	// receiving finishes first
	ch := make(chan *T, 5) 


	for i := range arr {
		count++;
		go send(i, &count, ch)
	}
	time.Sleep(time.Second)
	for i := range arr {
		count++;
		arr[i] = *<-ch
	}
}
