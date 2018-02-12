package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // HL
	case "darwin": // HL
		fmt.Println("OS X.")
	case "linux": // HL
		fmt.Println("Linux.")
	default: // HL
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.", os)
	} // HL
}
