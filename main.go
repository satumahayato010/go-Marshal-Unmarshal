package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"-"`
	Publisher string `json:"publisher"`
}

func (b Book) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Publisher string
	}{
		Publisher: b.Publisher + " Japan",
	})
	return v, err
}

func main() {
	b := []byte(`{"title": "リーダブルコード", "author": "Trevor Foucher", "Publisher": "OREILLY"}`)
	var book Book
	if err := json.Unmarshal(b, &book); err != nil {
		fmt.Println(err)
	}
	fmt.Println(book.Title, book.Author, book.Publisher)

	v, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(v))
}
