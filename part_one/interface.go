package partone

// Project ...
type Project interface {
	Print(string) string
	GetName() string
	SetName(string) string
}
