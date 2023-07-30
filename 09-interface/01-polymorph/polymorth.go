package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"os"
)

type serializer interface {
	serialize() ([]byte, error)
}

type person struct {
	Name string
	Age  int
}

func (p *person) serialize() ([]byte, error) {
	return json.Marshal(p)
}

type Car struct {
	Year int `json:"year"`
}

func (c *Car) serialize() ([]byte, error) {
	return xml.Marshal(c)
}

func store(data serializer, source io.WriteCloser) error {
	defer source.Close()
	b, err := data.serialize()
	if err != nil {
		return err
	}

	_, err = source.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var s serializer

	f, err := os.Create("./output.json")
	if err != nil {
		log.Fatal(err)
	}
	s = &person{}

	err = store(s, f)
	if err != nil {
		log.Fatal(err)
	}

}
