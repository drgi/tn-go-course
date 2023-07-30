package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func Test_sortInts(t *testing.T) {
	got := []int{1, 2, 3, 4}
	want := []int{1, 2, 3, 4}
	sort.Ints(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("sortInts() = %v, want %v", got, want)
	}
}

func Test_sortStrings(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test - 1. Empty slice",
			args: args{
				s: []string{},
			},
			want: []string{},
		},
		{
			name: "Test - 2. Simple eng chars",
			args: args{
				s: []string{"c", "b", "a"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test - 3. Simple eng and russian chars",
			args: args{
				s: []string{"c", "b", "a", "й", "б", "а", "я"},
			},
			want: []string{"a", "b", "c", "а", "б", "й", "я"},
		},
		{
			name: "Test - 4. Words eng and russian",
			args: args{
				s: []string{"World", "Hello", "Gopher", "Мир", "Привет", "Гофер", "Магия"},
			},
			want: []string{"Gopher", "Hello", "World", "Гофер", "Магия", "Мир", "Привет"},
		},
		{
			name: "Test - 5. Words and simbols eng and russian chars",
			args: args{
				s: []string{"World", "Hello", "Gopher", "Мир", "Привет", "Гофер", "Магия", "!", "@tag", "#хештег)"},
			},
			want: []string{"!", "#хештег)", "@tag", "Gopher", "Hello", "World", "Гофер", "Магия", "Мир", "Привет"},
		},
		{
			name: "Test - 7. Nil arg",
			args: args{
				s: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sort.Strings(tt.args.s); !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("sortStrings() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

// сравним сортировку int и float64 на количестве:
const sliceLength = 1_000_000

func BenchmarkSortInts(b *testing.B) {
	data := sampleIntSlice(sliceLength)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}

func BenchmarkSortFloat64(b *testing.B) {
	data := sampleFloatSlice(sliceLength)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Float64s(data)
	}
}

func sampleFloatSlice(length int) []float64 {
	s := make([]float64, length, length)
	for i := 0; i < length; i++ {
		s[i] = rand.ExpFloat64()
	}
	return s
}

func sampleIntSlice(length int) []int {
	s := make([]int, length, length)
	for i := 0; i < length; i++ {
		s[i] = rand.Int()
	}
	return s
}
