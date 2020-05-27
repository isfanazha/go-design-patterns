package persistance

import (
	"io/ioutil"
	"strings"

	"github.com/isfanazha/go-design-patterns/solid/singleresponsibility/domain"
)

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *domain.Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.Entries, p.LineSeparator)), 0644)
}

func (p *Persistence) LoadFile() {
	// ...
}
