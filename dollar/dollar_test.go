package dollar_test

import (
	"app/dollar"
	"testing"
)

/**
 * TODO
 *
 * - [ ] $5 + 10CHF = $10(レートが2:1の場合)
 * - [ ] amountをprivateにする
 * - [ ] Moneyの丸め処理をどうする?
 * - [ ] hasCode()
 * - [ ] nullとの等価性比較
 * - [ ] 他のオブジェクトとの等価性比較
 *
 * - [x] Dollarの副作用をどうする?
 * - [x] $5 * 2 = $10
 * - [x] equals()
 */
func TestMultiplication(t *testing.T) {
	five := dollar.Dollar{
		Amount: 5,
	}

	var product dollar.Dollar = five.Times(2)

	if product.Amount != 10 {
		t.Errorf("Actual %v must be 10", product.Amount)
	}

	product = five.Times(3)

	if product.Amount != 15 {
		t.Errorf("Actual %v must be 15", product.Amount)
	}
}

func TestEquality(t *testing.T) {
	five := dollar.Dollar{Amount: 5}

	var product dollar.Dollar = dollar.Dollar{Amount: 5}

	if !five.Equals(product) {
		t.Errorf("Actual %v must be 5", product.Amount)
	}

	product = dollar.Dollar{Amount: 6}

	if five.Equals(product) {
		t.Errorf("Actual %v must be 5", product.Amount)
	}
}
