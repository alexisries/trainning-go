package main

import "fmt"

type Instrumenter interface {
	Play()
}

type Amplifier interface {
	Connect(amp string)
}

type Guitar struct {
	Strings int
}

func (g Guitar) Play() {
	fmt.Printf("Tzzzzouiiiing with %d strings\n", g.Strings)
}

func (g Guitar) Connect(amp string) {
	fmt.Printf("Connected to %v\n", amp)
}

type Piano struct {
	Keys int
}

func (p Piano) Play() {
	fmt.Printf("Plip plip %d keys\n", p.Keys)
}

func main() {
	var instr Instrumenter
	instr = &Guitar{5}
	instr.Play()

	instr = &Piano{88}
	instr.Play()

	g := Guitar{12}
	g.Play()
	g.Connect("Marshall")
}
