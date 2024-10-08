package dollar_test

import (
	"app/dollar"
	"testing"
)

/**
 * TODO
 *
 * - [ ] $5 + 10CHF = $10(レートが2:1の場合)
 * - [x] $5 * 2 = $10
 * - [ ] amountをprivateにする
 * - [ ] Dollarの副作用をどうする?
 * - [ ] Moneyの丸め処理をどうする?
 */
func TestDollar(t *testing.T) {
	d := dollar.Dollar{
		Amount: 5,
	}
	d.Times(2)

	actual := d.Amount
	expect := 10

	if actual != expect {
		t.Errorf("Actual %v must be expect %v", actual, expect)
	}
}
