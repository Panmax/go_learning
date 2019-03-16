package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Pet1 struct {
}

func (p *Pet1) Speak() {
	fmt.Print("???")
}

func (p *Pet1) SpeakTo1(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
	Pet1
}

func (d *Dog) Speak() {
	fmt.Print("wang!")
}

func TestDog(t *testing.T) {
	//var dog *Pet = new(Dog)
	dog := new(Dog)
	dog.SpeakTo1("Pan")
}
