package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomthing(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if i, ok := p.(string); ok {
	//	fmt.Println("string", i)
	//	return
	//}
	//fmt.Println("Unknow Type")

	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomthing(1)
	DoSomthing("1")
	DoSomthing(1.3)
}
