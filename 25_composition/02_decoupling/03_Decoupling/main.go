package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	Line string
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

type PullStorer interface {
	Puller
	Storer
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
	Puller
	Storer
}

func Pull(p Puller, data []Data) (int, error) {
	for i := range data {
		err := p.Pull(&data[i])
		if err != nil {
			return i, err
		}
	}
	return len(data), nil
}
func Store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Copy(ps PullStorer, buf int) error {
	var data = make([]Data, buf)

	for {
		i, err := Pull(ps, data)
		if i > 0 {
			if _, err = Store(ps, data[:i]); err != nil {
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
		Puller: &Xenia{},  // Set db params
		Storer: &Pillar{}, // Set db params
	}
	Copy(&s, 3)

	x := s.Puller.(*Xenia) // Raising a flag
	fmt.Println("Total connections: ", x.Connections)
}
