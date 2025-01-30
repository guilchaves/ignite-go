package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	str := "🦆 quack quack!"
	reader := strings.NewReader(str)
	writer := MyWriter{os.Stdout}

	buffer := make([]byte, 1000)
	n, err := reader.Read(buffer)

	if err != nil {
		panic(err)
	}

	writer.Write(buffer[:n])
}

type MyWriter struct {
	w io.Writer
}

func (mw MyWriter) Write(b []byte) (int, error) {
	for i, bb := range b {
		b[i] = bb + 10
	}
	return mw.w.Write(b)
}
