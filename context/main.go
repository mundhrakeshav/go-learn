package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func first(ctx context.Context, list []string) (*result, error) {
	// results := make(chan result) //! Unbuffered channel,  This can't write till it has a listener
	results := make(chan result, len(list)) // Buffered channel can buffer values
	ctx, cancel := context.WithCancel(ctx)  // Attach a cancel
	defer cancel()
	for _, url := range list {
		go get(ctx, url, results)
	}
	select {
		case r := <-results:
			{
				return &r, nil // Return the first val received
			}
			case <-ctx.Done(): // In case parent context had a timeout
			{
				return nil, ctx.Err()
			}
		}
	}
	
	func get(ctx context.Context, url string, ch chan<- result) {
		var r result
		start := time.Now()
		ticker := time.NewTicker(1 * time.Second).C
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if resp, err := http.DefaultClient.Do(req); err != nil {
			t := time.Since(start).Round(time.Millisecond)
			r = result{url, err, t}
			} else {
				t := time.Since(start).Round(time.Millisecond)
				r = result{url, nil, t}
				resp.Body.Close()
			}
			
			for {
				select {
				case ch <- r:
					return
				//Leak: When the program sleeps and channel is unbuffered this 
				// 		case will be called again and again. As the channel 
				// 		being unbuffered won't be written to and be blocking.
			case <-ticker:
					log.Println("Tick", r)
				}
			}
		}
		
func main() {
	list := []string{
		"https://google.com",
		"https://amazon.com",
		"https://yahoo.com",
		"https://nytimes.com",
		"https://wsj.com",
		"https://gmail.com",
		"https://hih.com",
	}

	r, _ := first(context.Background(), list)

	if r.err != nil {
		fmt.Println("❌", r.err, r.latency)
	} else {
		fmt.Println("✅", r, r.latency.Seconds())
	}
	time.Sleep(time.Duration(time.Second * 9))

	fmt.Println(runtime.NumGoroutine())
}
