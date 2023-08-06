package main

import (
	"fmt"
	"io"
	"os"
)

type User interface {
	Age() uint
}

type Employee struct {
	name       string
	lastName   string
	speciality string
	age        uint
}

func (e *Employee) Age() uint {
	return e.age
}

func (e *Employee) FullName() string {
	return fmt.Sprintf("%s %s", e.lastName, e.name)
}

type Customer struct {
	name     string
	lastName string
	age      uint
}

func (c *Customer) Age() uint {
	return c.age
}

func main() {
	args := []interface{}{"В", 1, " ", true, "GO", false, " ", 'r', "нет", []byte{}, " ", &Customer{}, "магии", map[string]interface{}{}, " ", [2]int{}, 3.14, "!"}
	printString(os.Stdout, args...)
}

// вызов функции в тесте
func maxAge(users ...User) (max uint) {
	for _, u := range users {
		age := u.Age()
		if age > max {
			max = age
		}
	}
	return
}

func printString(w io.Writer, args ...interface{}) {
	for _, v := range args {
		switch t := v.(type) {
		case string:
			w.Write([]byte(t))
		default:
			continue
		}
	}
}
