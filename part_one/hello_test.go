package partone_test

import (
	"fmt"

	"github.com/wuxiaoxiaoshen/go-how-to-wirte-doc/part_one"
)

// ExampleNewHello is a example for Hello
func ExampleNewHello() {
	result := partone.NewHello("XieWei")
	fmt.Println(result.GetName())
}

func ExampleHello() {
	partone.HelloWorld()
}

func Example() {
	partone.HelloWorld()
	var example partone.Hello
	example.Name = "Golang doc"

	fmt.Println(example.GetName())
	fmt.Println(example.SetName("Python"))
}
