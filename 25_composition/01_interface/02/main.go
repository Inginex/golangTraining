package main

import (
	"fmt"

	"github.com/sampalm/07_compare/01/teste"
)

type Data struct {
	Chunk []byte
	Size  int64
}

func (d *Data) Print() {
	fmt.Println(string(d.Chunk))
}

type Printer interface {
	Print()
}

func PrintChunk(p Printer) {
	p.Print()
}

func main() {
	d := Data{Chunk: []byte("some-thing")}
	dt := teste.Data{Chunk: []byte("test-some-thing")}
	d.Print()
	dt.Print()
	fmt.Println("*********** INTERFACE **********")
	PrintChunk(&d)
	PrintChunk(&dt)
}
