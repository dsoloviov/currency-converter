package main

import (
	"testing"
)

const VALUE = 100
const EUR = "eur"
const USD = "usd"
const UAH = "uah"

func TestUsdToUah(t *testing.T) {
	c := Currencies{Usd: Currency{InverseRate: 0.5}, Uah: Currency{Rate: 4.0}}

	result := convert(USD, UAH, VALUE, c)
	if result != 200.0 {
		t.Errorf("Result was incorrect, got: %f, want: %f.", result, 200.0)
	}
}

func TestUahToEur(t *testing.T) {
	c := Currencies{Uah: Currency{InverseRate: 0.5}}

	result := convert(UAH, EUR, VALUE, c)
	if result != 50.0 {
		t.Errorf("Result was incorrect, got: %f, want: %f.", result, 50.0)
	}
}

func TestEurToUah(t *testing.T) {
	c := Currencies{Uah: Currency{Rate: 2.0}}

	result := convert(EUR, UAH, VALUE, c)
	if result != 200 {
		t.Errorf("Result was incorrect, got: %f, want: %f.", result, 10.0)
	}
}
