package main

import "fmt"

type ReadCloser interface {
	Read([]byte) (int, error)
	Close() error
}

type ByteReader byte

func (r ByteReader) Close() error { return nil }

func (r ByteReader) Read(p []byte) (int, error) {
	if len(p) < 1 {
		return 0, fmt.Errorf("no space!")
	}
	p[0] = byte(r)
	return 1, nil
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

func main() {
	b := make([]byte, 15)
	r := ByteReader(135)
	ReadAndClose(r, b)
	fmt.Println(b)
}
