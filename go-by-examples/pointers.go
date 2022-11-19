package main

import "fmt"

// pass by value
func zeroval(ival int) {
    ival = 0
}

// pass by reference
func zeroptr(iptr *int) {
    *iptr = 12
}

func pointers() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}