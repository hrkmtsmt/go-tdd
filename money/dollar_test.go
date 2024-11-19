package money_test

import (
	m "app/money"
	"testing"
)

/**
 * TODO
 *
 * - [ ] equalsの一般化
 *
 * - [ ] DollarとFrancの重複
 * - [ ] timesの一般化
 * - [ ] $5 + 10CHF = $10(レートが2:1の場合)
 * - [ ] Moneyの丸め処理をどうする?
 * - [ ] hasCode()
 * - [ ] nullとの等価性比較
 * - [ ] 他のオブジェクトとの等価性比較
 * - [ ] FrancとDollarを比較する
 *
 * - [x] Dollarの副作用をどうする?
 * - [x] $5 * 2 = $10
 * - [x] equals()
 * - [x] amountをprivateにする
 * - [x] 5CHF * 2 = 10CHF
 */
func TestDollarMultiplication(t *testing.T) {
	five := m.NewDollar(5)

	ten := m.NewDollar(10)

	if ten != five.Times(2) {
		t.Errorf("Actual %v must be %v", ten, five.Times(2))
	}

	fifteen := m.NewDollar(15)

	if fifteen != five.Times(3) {
		t.Errorf("Actual %v must be %v", fifteen, five.Times(3))
	}
}

func TestDollarEquality(t *testing.T) {
	five := m.NewDollar(5)

	var product m.Dollar = m.NewDollar(5)

	if !five.Equals(product) {
		t.Errorf("Actual %v must be 5", product)
	}

	product = m.NewDollar(6)

	if five.Equals(product) {
		t.Errorf("Actual %v must be 5", product)
	}
}
