package main

import "fmt"

type StringIntMap struct {
	Map map[string]int
}

func (m *StringIntMap) Add(key string, value int) {
	m.Map[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.Map, key)
}

func (m *StringIntMap) Copy() map[string]int {
	new := make(map[string]int)
	for k, v := range m.Map {
		new[k] = v
	}

	return new
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.Map[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	v, ok := m.Map[key]
	return v, ok
}

func main() {
	m := StringIntMap{
		Map: make(map[string]int),
	}

	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)
	m.Add("z", 3)
	fmt.Printf("Created: %v\n", m)

	m.Remove("c")
	fmt.Printf("Element removed: %v\n", m)

	copyMap := m.Copy()
	fmt.Printf("Copy: %v\n", copyMap)

	fmt.Printf("Checking 'a' and 'd' existence: %v, %v\n", m.Exists("a"), m.Exists("d"))
}
