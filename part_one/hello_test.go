package partone_test

import (
	"fmt"
	"testing"

	partone "github.com/wuxiaoxiaoshen/go-how-to-wirte-doc/part_one"
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

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			partone.HelloWorld()
		})
	}
}
