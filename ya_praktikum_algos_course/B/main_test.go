package main

import "testing"

func GeneralCheck(t *testing.T, input string, expected string) {
	got := ReverseNumber(input)
	if got != expected {
		t.Errorf("ReverseNumber(%s) = %s; want %s", input, got, expected)
	}
}

func TestReverseNumber123(t *testing.T) {
	GeneralCheck(t, "123", "321")
}

func TestReverseNumberMinus150(t *testing.T) {
	GeneralCheck(t, "-150", "-51")
}

func TestReverseNumberZero(t *testing.T) {
	GeneralCheck(t, "0", "0")
}

func TestReverseNumberManyZeros(t *testing.T) {
	GeneralCheck(t, "0000", "0")
}

func TestReverseNumberMinusOneThousandFive(t *testing.T) {
	GeneralCheck(t, "-1005", "-5001")
}

func TestReverseNumberMinusOneHundredFifty(t *testing.T) {
	GeneralCheck(t, "-1050", "-501")
}

func TestReverseNumberMinusOneHundred(t *testing.T) {
	GeneralCheck(t, "-1000", "-1")
}

func TestReverseNumberMinusOne(t *testing.T) {
	GeneralCheck(t, "-0001", "-1000")
}
