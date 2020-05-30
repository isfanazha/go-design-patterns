package usecase

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/solid/depedencyinversion/repository"
)

// Research is HLM (High Level Module)
// HLM should not depend on LLM (Low Level Module)
// Both should depend on abstractions
type Research struct {
	Browser repository.RelationshipBrowser // LLM

	// If we use this approach, Research as HLM depends on LLM (relationship).
	// relationships repository.Relationships
}

func (r *Research) Investigate() {
	for _, p := range r.Browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.Name)
	}

	// We can use this code below, if Research depends on LLM (relationship).
	//relations := r.relationships.relations
	//for _, rel := range relations {
	//	if rel.from.name == "John" &&
	//		rel.relationship == Parent {
	//		fmt.Println("John has a child called", rel.to.name)
	//	}
	//}

	// Imagine if we change relationships.relations to another type, we need to adjust the code.
	// So we had better use abstraction/interface, otherwise we must to change/adjust the code whenever there is change.
}
