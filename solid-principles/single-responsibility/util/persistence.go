package util

import (
	"io/ioutil"
	"strings"

	"github.com/isfanazha/go-design-patterns/solid-principles/single-responsibility/domain"
)

var lineSeparator = "\n"

func SaveToFile(j *domain.Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.Entries, lineSeparator)), 0644)
}

func LoadFromFile() {
	// ...
}
