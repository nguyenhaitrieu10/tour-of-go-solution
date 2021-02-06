// Exercise: rot13Reader
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i, v := range b {
		if v > 'z' || v < 'A' {
			continue
		}
		if v < 'a' {
			b[i] = (v-'A'+13)%26 + 'A'
		} else {
			b[i] = (v-'a'+13)%26 + 'a'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
