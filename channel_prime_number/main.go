package main

func generator(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		ch <- i;
	}
	close(ch)
}

/// @param src: channel to read from
/// @param dst: channel to write to
/// @param prime: prime number it filters on
func filter(src <-chan int, dst chan <-int, prime int)  {
	for i := range src {
		if i % prime != 0 {
			dst <- i
		}
	}
	close(dst)
}

func seive(limit int)  {
	ch := make(chan int) // channel bw main and generator
	go generator(limit, ch);

	for {
		//Channel gives 2 values 1st is the actual value and second is a bool which tells if the channel is open or closed
		prime, ok := <-ch
		if !ok {
			break
		}

		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1;
		// fmt.Print(prime, " ")
	}
	
}

func main() {
	seive(1000000)
}
