package domain

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)
