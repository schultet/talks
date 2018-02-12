package main

import "fmt"

type ReadCloser interface {
	Read([]byte) (int, error)
	Close()
}

func ReadAndClose(r ReadCloser, b []byte) (int, error) { // HL
	var n, nr int
	var err error
	for len(b) > 0 && err == nil {
		nr, err = r.Read(b) // HL
		n += nr
		b = b[nr:]
	}
	r.Close() // HL
	return n, err
}

type MyByte byte

func (rc MyByte) Read(b []byte) (int, error) { // HL
	if len(b) == 0 {
		return 0, nil
	}
	b[0] = byte(rc)
	return 1, nil
}

func (rc MyByte) Close() {} // HL

func main() {
	mb := MyByte(35)
	b := make([]byte, 10)
	ReadAndClose(mb, b) // HL
	fmt.Println(b)
}
