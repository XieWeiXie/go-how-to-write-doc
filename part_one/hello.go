package partone

import "fmt"

// project info
const (
	Version     = "0.0.1"                        //  version for project
	ProjectName = "How to write doc with golang" // project name
)

// name
var (
	Name = "XieWei" // name for one person
)

// HelloWorld is one function just print: Hello World
// Deprecated: should be deprecated because of no use
func HelloWorld() {
	fmt.Println("Hello World")
}

// Hello is a struct with one field: Name
type Hello struct {
	Name string `json:"name"`
}

// NewHello is a initializer for Hello
func NewHello(value string) *Hello {
	return &Hello{
		Name: value,
	}
}

// Print will return a string.
func (h Hello) Print(value string) string {
	return fmt.Sprintf("%s:%s", h.Name, value)
}

// GetName will return struct Hello field name.
func (h Hello) GetName() string {
	return h.Name
}

// SetName will set a name for struct Hello.
func (h *Hello) SetName(value string) string {
	h.Name = value
	return h.Name
}
