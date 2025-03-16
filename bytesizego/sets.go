package main

// Set is a collection of unique elements
type Set struct {
	elements map[string]struct{}
}

// NewSet creates a new set
func NewSet() *Set {
	return &Set{
		elements: make(map[string]struct{}),
	}
}

// Add inserts an element into the set
func (s *Set) Add(value string) {
	s.elements[value] = struct{}{}
}

// Remove deletes an element from the set
func (s *Set) Remove(value string) {
	delete(s.elements, value)
}

// Contains checks if an element is in the set
func (s *Set) Contains(value string) bool {
	_, found := s.elements[value]
	return found
}

// Size returns the number of elements in the sets
func (s *Set) Size() int {
	return len(s.elements)
}

// List returns all elements in the set as a slice
func (s *Set) List() []string {
	keys := make([]string, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	set := NewSet()

	// Add elements to the set
	set.Add("apple")
	set.Add("banana")
	set.Add("orange")

}
