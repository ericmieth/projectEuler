package main

import (
	"testing"
)

func TestGetNextHorizontally(t *testing.T) {
	cases := []struct {
		input []coord
		want  int
	}{
		{[]coord{coord{1, 1, 4}, coord{2, 1, 5}}, 5},
		{[]coord{coord{2, 1, 4}, coord{3, 1, 7}}, 7},
	}
	for _, c := range cases {
		got := c.input[0].getNextHorizontally(1, c.input)
		if got != c.want {
			t.Errorf("getNextHorizontally(%v) ==\n%v (is)\n%v (should)", c.input, got, c.want)
		}
	}
}

func TestGetNextVerically(t *testing.T) {
	cases := []struct {
		input []coord
		want  int
	}{
		{[]coord{coord{1, 1, 4}, coord{1, 2, 5}}, 5},
		{[]coord{coord{1, 3, 4}, coord{1, 4, 7}}, 7},
	}
	for _, c := range cases {
		got := c.input[0].getNextVertically(1, c.input)
		if got != c.want {
			t.Errorf("getNextVertically(%v) ==\n%v (is)\n%v (should)", c.input, got, c.want)
		}
	}
}

func TestGetNextDiagonallyNWtoSE(t *testing.T) {
	cases := []struct {
		input []coord
		want  int
	}{
		{[]coord{coord{1, 1, 4}, coord{2, 2, 5}}, 5},
		{[]coord{coord{4, 4, 4}, coord{5, 5, 7}}, 7},
	}
	for _, c := range cases {
		got := c.input[0].getNextDiagonallyNWtoSE(1, c.input)
		if got != c.want {
			t.Errorf("getNextDiagonallyNWtoSE(%v) ==\n%v (is)\n%v (should)", c.input, got, c.want)
		}
	}
}

func TestGetNextDiagonallySWtoNE(t *testing.T) {
	cases := []struct {
		input []coord
		want  int
	}{
		{[]coord{coord{1, 2, 4}, coord{2, 1, 5}}, 5},
		{[]coord{coord{4, 4, 4}, coord{5, 3, 7}}, 7},
	}
	for _, c := range cases {
		got := c.input[0].getNextDiagonallySWtoNE(1, c.input)
		if got != c.want {
			t.Errorf("getNextDiagonallySWtoNE(%v) ==\n%v (is)\n%v (should)", c.input, got, c.want)
		}
	}
}

func TestMultiplyHorizontally(t *testing.T) {
	cases := []struct {
		input []coord
		want  int
	}{
		{[]coord{coord{1, 1, 8}, coord{2, 2, 2}, coord{3, 2, 22}}, 352},
		{[]coord{coord{4, 4, 4}, coord{5, 5, 7}, coord{6, 5, 2}}, 56},
		{[]coord{coord{4, 4, 2}, coord{5, 5, 7}, coord{6, 5, 2}}, 28},
		{[]coord{coord{4, 4, 3}, coord{5, 5, 7}, coord{6, 5, 0}}, 0},
	}
	for _, c := range cases {
		got := c.input[0].multilyHorizontally(3, c.input)
		if got != c.want {
			t.Errorf("multilyHorizontally(%v) ==\n%v (is)\n%v (should)", c.input, got, c.want)
		}
	}
}

func TestMultiplyVertically(t *testing.T) {
	cases := []struct {
		input []coord
		want  int
	}{
		{[]coord{coord{1, 1, 8}, coord{1, 2, 2}, coord{1, 3, 22}}, 352},
		{[]coord{coord{4, 4, 4}, coord{4, 5, 7}, coord{4, 6, 2}}, 56},
		{[]coord{coord{2, 4, 2}, coord{2, 5, 7}, coord{2, 6, 2}}, 28},
		{[]coord{coord{5, 4, 3}, coord{5, 5, 7}, coord{5, 6, 0}}, 0},
	}
	for _, c := range cases {
		got := c.input[0].multilyVertically(3, c.input)
		if got != c.want {
			t.Errorf("multilyVertically(%v) ==\n%v (is)\n%v (should)", c.input, got, c.want)
		}
	}
}

func TestMultiplyDiagonallyNWtoSE(t *testing.T) {
	cases := []struct {
		input []coord
		want  int
	}{
		{[]coord{coord{1, 1, 8}, coord{2, 2, 2}, coord{3, 3, 22}}, 352},
		{[]coord{coord{3, 4, 4}, coord{4, 5, 7}, coord{5, 6, 2}}, 56},
		{[]coord{coord{1, 4, 2}, coord{2, 5, 7}, coord{3, 6, 2}}, 28},
		{[]coord{coord{5, 6, 3}, coord{5, 6, 7}, coord{7, 7, 0}}, 0},
	}
	for _, c := range cases {
		got := c.input[0].multilyDiagonallyNWtoSE(3, c.input)
		if got != c.want {
			t.Errorf("multilyDiagonallyNWtoSE(%v) ==\n%v (is)\n%v (should)", c.input, got, c.want)
		}
	}
}

func TestMultiplyDiagonallySWtoNE(t *testing.T) {
	cases := []struct {
		input []coord
		want  int
	}{
		{[]coord{coord{3, 3, 8}, coord{4, 2, 2}, coord{5, 1, 22}}, 352},
	}
	for _, c := range cases {
		got := c.input[0].multilyDiagonallySWtoNE(3, c.input)
		if got != c.want {
			t.Errorf("multilyDiagonallySWtoNE(%v) ==\n%v (is)\n%v (should)", c.input, got, c.want)
		}
	}
}
