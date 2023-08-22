package intface

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_maxAge(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name    string
		args    args
		wantMax uint
	}{
		{
			name: "Test - 1, valid users",
			args: args{users: []User{
				&Employee{
					name:       "Иван",
					lastName:   "Петров",
					age:        25,
					speciality: "frontEnd",
				},
				&Customer{
					name:     "Петр",
					lastName: "Иванов",
					age:      37,
				},
				&Employee{
					name:       "Федор",
					lastName:   "Инженеров",
					age:        25,
					speciality: "backEnd",
				},
				&Customer{
					name:     "Сергей",
					lastName: "Иванов",
					age:      60,
				},
			}},
			wantMax: 60,
		},
		{
			name:    "Test - 2, empty slice",
			args:    args{users: []User{}},
			wantMax: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxAge(tt.args.users...); gotMax != tt.wantMax {
				t.Errorf("maxAge() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func Test_printString(t *testing.T) {
	args := []interface{}{"В", 1, " ", true, "GO", false, " ", 'r', "нет", []byte{}, " ", &Customer{}, "магии", map[string]interface{}{}, " ", [2]int{}, 3.14, "!"}
	w := &bytes.Buffer{}
	printString(w, args...)
	got := w.String()
	want := "В GO нет магии !"
	if got != want {
		t.Errorf("printString() = %v, want %v", got, want)
	}
}

func Test_maxAgePerson(t *testing.T) {
	type args struct {
		users []interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult interface{}
	}{
		{
			name: "Test - 1, valid users",
			args: args{users: []interface{}{
				&Employee{
					name:       "Иван",
					lastName:   "Петров",
					age:        25,
					speciality: "frontEnd",
				},
				&Customer{
					name:     "Петр",
					lastName: "Иванов",
					age:      37,
				},
				&Employee{
					name:       "Федор",
					lastName:   "Инженеров",
					age:        25,
					speciality: "backEnd",
				},
				&Customer{
					name:     "Сергей",
					lastName: "Иванов",
					age:      60,
				},
			}},
			wantResult: &Customer{
				name:     "Сергей",
				lastName: "Иванов",
				age:      60,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := maxAgePerson(tt.args.users...); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("maxAgePerson() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
