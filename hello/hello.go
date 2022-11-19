package main

import (
	"fmt"
)

func main() {
	m := 12;
var x *int = &m;
z := &x;
fmt.Println(z);
}
