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
 *
 * - [x] Dollarの副作用をどうする?
 * - [x] $5 * 2 = $10
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
