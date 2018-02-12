package main

import (
	"fmt"
	"sort"
)

type ByteSlice []byte

// TODO: implement methods of sort.Interface
// func (b ByteSlice) Len() int { ... }
// ...

func main() {
	b := ByteSlice{2, 0, 133, 44, 10, 200}
	sort.Sort(b)
	fmt.Println(b)
	// Output: [0 2 10 44 133 200]
}
