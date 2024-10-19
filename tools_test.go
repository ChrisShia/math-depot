package maths

import (
	"fmt"
	"testing"
)

func Test_DefaultCounterModulo(t *testing.T) {
	tests := []struct {
		name     string
		advance  int
		modulo   int
		expected int
	}{
		{"Start from 0 advance 5 times mod 3", 5, 3, 1},
	}
	for _, tt := range tests {
		counter := DefaultModCounter(tt.modulo)
		t.Run(tt.name, func(t *testing.T) {
			for i := 1; i <= tt.advance; i++ {
				c := counter.ModIndex()
				fmt.Println(c)
				counter.Next()
			}
		})
	}
}

func Test_CounterModulo(t *testing.T) {
	tests := []struct {
		name     string
		advance  int
		start    int
		modulo   int
		expected int
	}{
		{"Start from -2 advance 5 times mod 3", 5, -2, 3, 0},
		{"Start from 100 advance 5 times mod 3", 5, 100, 3, 102},
	}
	for _, tt := range tests {
		counter := NewModCounter(tt.start, tt.modulo)
		t.Run(tt.name, func(t *testing.T) {
			for i := 1; i <= tt.advance; i++ {
				counter.Next()
			}
			ans := counter.ModIndex()
			if ans != tt.expected {
				t.Errorf("Actual = %d, expected %d", ans, tt.expected)
			}
		})
	}
}
