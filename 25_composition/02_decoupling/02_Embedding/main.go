package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	Line string
}

type Xenia struct {
	Connections int
}

func (x *Xenia) Pull(d *Data) error {
	r := rand.Intn(10)
	if r == 5 {
		return fmt.Errorf("Failed to connect")
	}
	x.Connections++
	d.Line = "New line."
	return nil
}

type Pillar struct {
	Stored int
}

func (p *Pillar) Store(d *Data) error {
	r := rand.Intn(6)
	if r == 2 {
		return fmt.Errorf("Failed to store")
	}
	p.Stored++
	fmt.Println("Stored: ", d.Line)
	return nil
}

type System struct {
	Xenia
	Pillar
}

func Pull(x *Xenia, data []Data) (int, error) {
	for i := range data {
		err := x.Pull(&data[i])
		if err != nil {
			return i, err
		}
	}
	return len(data), nil
}
func Store(p *Pillar, data []Data) (int, error) {
	for i := range data {
		if err := p.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Copy(s *System, buf int) error {
	var data = make([]Data, buf)

	for {
		i, err := Pull(&s.Xenia, data)
		if i > 0 {
			if _, err = Store(&s.Pillar, data[:i]); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	s := System{
		Xenia:  Xenia{},  // Set db params
		Pillar: Pillar{}, // Set db params
	}
	Copy(&s, 3)
	fmt.Printf("Connections: %d - Stored Data: %d\n", s.Connections, s.Stored)
}
