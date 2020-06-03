package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (input rot13Reader) Read(out []byte)(int,error){
	chars := len(out)
	in := make([]byte,chars)
	inputLength, err := input.r.Read(in)
	if err !=nil{
		return 0,err
	}
	for i:=0;i<inputLength;i++{
		char := in[i]
		if !unicode.IsLetter(rune(char)){
			out[i] = char
			continue
		}
		// preChar 01x  00001 alphabet
		preChar := char & byte(224)
		alphabet := char & byte(31)
		alphabet += 13
		alphabet %= 26
		char = preChar + alphabet
		out[i] = char
	}
	return inputLength,nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
