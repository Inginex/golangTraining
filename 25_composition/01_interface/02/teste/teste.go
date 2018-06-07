package teste

import (
	"fmt"
)

type Data struct {
	Chunk []byte
	Size  int64
}

func (d *Data) Print() {
	fmt.Println(string(d.Chunk))
}
