package money_test

import (
	m "app/money"
	"testing"
)

func TestFrancMultiplation(t *testing.T) {
	five := m.NewFranc(5)

	ten := m.NewFranc(10)

	if ten != five.Times(2) {
		t.Errorf("Actual %v must be %v", ten, five.Times(2))
	}

	fifteen := m.NewFranc(15)

	if fifteen != five.Times(3) {
		t.Errorf("Actual %v must be %v", fifteen, five.Times(3))
	}
}

func TestFrancEquality(t *testing.T) {
	five := m.NewFranc(5)

	var product m.Franc = m.NewFranc(5)

	if !five.Equals(product) {
		t.Errorf("Actual %v must be 5", product)
	}

	product = m.NewFranc(6)

	if five.Equals(product) {
		t.Errorf("Actual %v must be 5", product)
	}
}
