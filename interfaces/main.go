package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	sq := square{sideLength: 10}
	tri := triangle{
		base:   5,
		height: 10,
	}
	printArea(sq)
	printArea(tri)
	printGreeting(eb)
	printGreeting(sb)
	readFileFromFileName()
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello!!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!!"
}
