package main

import (
	"testing"
)

func TestAddColumn(t *testing.T) {
	cases := []struct {
		inColumn         []string
		inCarry          int
		wantOutLastDigit string
		wantOutCarry     int
	}{
		{[]string{"3", "4", "5"}, 0, "2", 1},
		{[]string{"3", "4", "5"}, 2, "4", 1},
		{[]string{"9", "5", "9"}, 0, "3", 2},
	}

	for _, c := range cases {
		gotOutLastDigit, gotOutCarry := addColumn(c.inColumn, c.inCarry)
		if gotOutLastDigit != c.wantOutLastDigit || gotOutCarry != c.wantOutCarry {
			t.Errorf(
				"addColumn(%v,%v) ==\n%v\t%v (is)\n%v\t%v (should)",
				c.inColumn,
				c.inCarry,
				gotOutLastDigit,
				gotOutCarry,
				c.wantOutLastDigit,
				c.wantOutCarry)
		}
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"3", "4", "5"}, "543"},
		{[]string{"1", "2", "3", "4"}, "4321"},
	}

	for _, c := range cases {
		got := reverse(c.in)
		if got != c.want {
			t.Errorf(
				"reverse(%v) ==\n%v (is)\n%v (should)",
				c.in,
				got,
				c.want,
			)
		}
	}
}
