package intface

import (
	"fmt"
	"io"
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

// вызов функции в тесте
func maxAgePerson(users ...interface{}) (result interface{}) {
	var maxAge uint
	for _, u := range users {
		switch t := u.(type) {
		case *Employee:
			if t.age > maxAge {
				maxAge = t.age
				result = t
			}
		case *Customer:
			if t.age > maxAge {
				maxAge = t.age
				result = t
			}
		default:
			continue
		}
	}
	return
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
		if s, ok := v.(string); ok {
			w.Write([]byte(s))
		}
	}
}
