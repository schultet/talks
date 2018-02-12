package main

import "fmt"

type ReadCloser interface {
	Read([]byte) (int, error)
	Close() error
}

type ByteReader byte

// TODO: Add a Read([]byte) (int, error) method to ByteReader.
// TODO: Add a Close() error method to ByteReader.

func main() {
	b := make([]byte, 5)
	r := ByteReader(35)
	ReadAndClose(r, b)
	fmt.Println(b)
	// OUTPUT: [35 35 35 35 35]
}

func ReadAndClose(r ReadCloser, b []byte) (int, error) {
	nx, n := 0, 0
	var err error
	for len(b) > 0 && err == nil {
		n, err = r.Read(b)
		nx += n
		b = b[n:]
	}
	r.Close()
	return nx, err
}
