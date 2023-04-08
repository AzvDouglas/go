package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 100.0
	expected := 10.0

	actual := CalculateTax(amount)

	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}