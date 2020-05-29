package domain

type Color int

const (
	Red Color = iota
	Green
	Blue
)

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type Product struct {
	Name  string
	Color Color
	Size  Size
}
