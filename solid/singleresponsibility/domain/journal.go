package domain

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Journal has the following functions:
// 1. AddEntry: Add an entry in journal data
// 2. RemoveEntry: Remove an entry in journal data
// 3. String: Joining entries data
// 4. SaveToFile: Save journal data into a file (Need to move to different type/package)
// 5. LoadFile: Load journal data from a file (Need to move to different type/package)

type Journal struct {
	// Journal entries
	Entries []string
	// Count entry index
	EntryCount int
}

func (j *Journal) AddEntry(text string) int {
	j.EntryCount++
	entry := fmt.Sprintf("%d: %s", j.EntryCount, text)
	j.Entries = append(j.Entries, entry)
	return j.EntryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

func (j *Journal) String() string {
	return strings.Join(j.Entries, "\n")
}

// NOTE: We need to move Save() and Load() to different types/packages.
// Those methods have independent tasks/problems, so we need to move them into different package/type.
// In this example, we move it into persistence package as persistence struct type and util package as a function.

func (j *Journal) SaveToFile(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) LoadFile(filename string) {
	// ...
}
