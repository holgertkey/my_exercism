package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

func SayHello(name string, g Greeter) string {
    return g.Greet(name)
}

type Italian struct {

}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(name string) string {
    return "I can speak " + i.LanguageName() + ": Ciao " + name + "!" 
}

type Portuguese struct {

}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet (name string) string {
    return "I can speak " + p.LanguageName() + ": Olá " + name + "!"
}