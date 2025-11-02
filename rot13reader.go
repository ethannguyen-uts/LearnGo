package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (p *rot13Reader) Read(b []byte) (int, error) {
	n, err := p.r.Read(b)
	
	for i := range n {
		rep := getChar(rune(b[i]))
		b[i] = byte(rep)
    }
    return n, err
}

func getChar(char rune) rune {
	        if char >= 'a' && char <= 'z' {
			// Lowercase letters: Shift by 13, wrapping around
			char = char - 'a' // Convert 'a' to 0, 'b' to 1, etc.
			char = (char + 13) % 26 // Apply the shift and wrap with modulo
			char = char + 'a' // Convert back to the shifted lowercase letter
		} else if char >= 'A' && char <= 'Z' {
			// Uppercase letters: Shift by 13, wrapping around
			char = char - 'A' // Convert 'A' to 0, 'B' to 1, etc.
			char = (char + 13) % 26 // Apply the shift and wrap with modulo
			char = char + 'A' // Convert back to the shifted uppercase letter
		}
	return char
}

func cRot13Read() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
