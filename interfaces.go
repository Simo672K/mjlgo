package main

import "fmt"

type Storer interface {
	GetItems() ([]any, error)
	Put(item any) error
}

type DBStore struct {
	dbStore Storer
}

func (db DBStore) GetItems() ([]any, error) {
	return []any{1, 2, "a", "some other data"}, nil
}

func (db DBStore) Put(item any) error {
	store := []any{}
	store = append(store, item)

	return nil
}

type interfaceTest struct {
	storeTest Storer
}

func main() {
	test := interfaceTest{
		storeTest: DBStore{},
	}

	items, error := test.storeTest.GetItems()

	if error != nil {
		panic("There is an error")
	}

	for key, val := range items {
		fmt.Println("item index:", key, "with value:", val)
	}

}
