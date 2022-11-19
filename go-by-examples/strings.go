package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func str() {
  hi := "HI"

    // getting address as string dynamically to preserve compatibility
    address := fmt.Sprint(&hi)

    fmt.Printf("Address of var hi: %s\n", address)

    // convert to uintptr
    var adr uint64
    adr, err := strconv.ParseUint(address, 0, 64)
    if err != nil {
        panic(err)
    }
    var ptr uintptr = uintptr(adr)

	fmt.Println(int(*(*byte)(unsafe.Pointer((ptr)))))	
    // fmt.Printf("String at address: %s\n", address)
    // fmt.Println(ptrToString(ptr))
}


// func ptrToString(ptr uintptr) string {
//     p := unsafe.Pointer(ptr + 1)

//     // return *(*string)(p)
// }