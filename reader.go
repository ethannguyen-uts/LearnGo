package main

import (
	"fmt"
	"io"
	"strings"
)

func cRead() {
	 var r *strings.Reader = strings.NewReader("Hello world this is my reader")

	 var b []byte = make([]byte, 8)
	 for {
		n, err := r.Read(b)
		fmt.Printf("n = %d err = %v b = %v", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	 }
}