package main

import "code.google.com/p/go-tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	b = append(b, 'A')
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
