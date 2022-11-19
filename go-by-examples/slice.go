package main

import "fmt"

func slice() {

    s := make([]string, 0)
    fmt.Println("emp:", s)
    s = append(s, "Kess")
    fmt.Println("set:", s)
    s = append(s, "Kess")
    fmt.Println("set:", s)
    s = append(s, "Kess")

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s)

    fmt.Println("len:", len(s))

    s = append(s, "d")
    fmt.Println("apd:", s)
    s = append(s, "e")
    fmt.Println("apd:", s)
    s = append(s, "f")

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}


func slicePtr(slicePtr *[]int)  {
    fmt.Println(*slicePtr)
    slicex := *slicePtr;
    *slicePtr = slicex[0:len(slicex) - 1]
}