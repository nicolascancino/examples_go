package main

import "testing"

func TestSum(t *testing.T) {
	/*total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	}*/

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}

}

func TestGetMax(t *testing.T) {

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 2},
		{5, 4, 5},
	}

	for _, item := range tables {
		numMax := GetMax(item.a, item.b)
		if numMax != item.n {
			t.Errorf("Max number was incorrect, got %d expected %d", numMax, item.n)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Max number was incorrect, got %d expected %d", fib, item.n)
		}
	}
}
