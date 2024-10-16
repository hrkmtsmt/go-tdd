package dollar_test

import (
	"app/dollar"
	"testing"
)

/**
 * TODO
 *
 * - [ ] $5 + 10CHF = $10(レートが2:1の場合)
 * - [ ] Moneyの丸め処理をどうする?
 * - [ ] hasCode()
 * - [ ] nullとの等価性比較
 * - [ ] 他のオブジェクトとの等価性比較
 *
 * - [x] Dollarの副作用をどうする?
 * - [x] $5 * 2 = $10
 * - [x] equals()
 * - [x] amountをprivateにする
 */
func TestMultiplication(t *testing.T) {
	five := dollar.NewDollar(5)

	ten := dollar.NewDollar(10)

	if ten != five.Times(2) {
		t.Errorf("Actual %v must be %v", ten, five.Times(2))
	}

	fifteen := dollar.NewDollar(15)

	if fifteen != five.Times(3) {
		t.Errorf("Actual %v must be %v", fifteen, five.Times(3))
	}
}

func TestEquality(t *testing.T) {
	five := dollar.NewDollar(5)

	var product dollar.Dollar = dollar.NewDollar(5)

	if !five.Equals(product) {
		t.Errorf("Actual %v must be 5", product)
	}

	product = dollar.NewDollar(6)

	if five.Equals(product) {
		t.Errorf("Actual %v must be 5", product)
	}
}
