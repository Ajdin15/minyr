package conv

import (
	"reflect"
	"testing"
)

/*
*
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 134},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}

}

func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 329.82},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 329.82},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) ) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 56.67},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 134},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}