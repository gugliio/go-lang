package main

import "testing"

func TestSum(t *testing.T) {
	// total := Sum(5, 5)
	// if total != 10 {
	//	t.Errorf("Error en el test de suma. Tuvimos un resultado de %d se esperaba %d", total, 10)
	// }
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
			t.Errorf("Error en el test de suma. Tuvimos un resultado de %d se esperaba %d", total, item.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{7, 9, 9},
		{5, 4, 5},
		{25, 26, 26},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("Error en el test de Get Max. Tuvimos un resultado de %d se esperaba %d", max, item.n)
		}
	}
}
