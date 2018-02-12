package main

import "fmt"

func main() {
	orange := map[string]byte{"red": 0xff, "green": 0x99, "blue": 0x00}

	for k, v := range orange { // HL
		fmt.Printf("%v: %v\n", k, v)
	} // HL

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow { // HL
		fmt.Printf("2**%d = %d\n", i, v)
	} // HL
}
