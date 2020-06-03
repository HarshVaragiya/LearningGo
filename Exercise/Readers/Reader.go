package main

import (
	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	chars := len(b)
	for i:=0;i<chars;i++{
		b[i]= byte('A')
	}
	return chars,nil
}

func main() {
	reader.Validate(MyReader{})
	b := make([]byte,16)
	_, _ = MyReader{}.Read(b)
	fmt.Println(b)
}
