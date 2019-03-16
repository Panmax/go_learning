package polymorphism

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello World!\")"
}

func wirteFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	wirteFirstProgram(goProg)
	wirteFirstProgram(javaProg)
}
