package main

import (
	"fmt"
	"sort"
)

type ByteSlice []byte

func (b ByteSlice) Len() int           { return len(b) }
func (b ByteSlice) Less(i, j int) bool { return b[i] < b[j] }
func (b ByteSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func main() {
	b := ByteSlice{2, 0, 133, 44, 10, 200}
	sort.Sort(b)
	fmt.Println(b)
	// Output: [0 2 10 44 133 200]
}
