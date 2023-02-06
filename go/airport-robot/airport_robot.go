package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(visitor string) string
}

func SayHello(visitor string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(visitor))
}

type German struct{}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(visitor string) string {
	return fmt.Sprintf("Hallo %s!", visitor)
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(visitor string) string {
	return fmt.Sprintf("Ciao %s!", visitor)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(visitor string) string {
	return fmt.Sprintf("Olá %s!", visitor)
}
