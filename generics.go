package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	bucket map[K]V
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		bucket: make(map[K]V),
	}
}

func (m *CustomMap[K, V]) Insert(key K, val V) error {
	m.bucket[key] = val
	return nil
}

func (m *CustomMap[K, V]) Get(key K) (output V, err error) {
	return m.bucket[key], nil
}

func Generics() {
	testMap := NewCustomMap[string, int]()
	testMap.Insert("Hello", 1)

	getHello, err := testMap.Get("Hello")
	if err != nil {
		panic("Error accured!")
	}

	fmt.Println("Hello keyword is mapped to: ", getHello)
}
