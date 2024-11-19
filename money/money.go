package money

type Money struct {
	amount int
}

func NewMoney(amount int) Money {
	return Money{amount}
}

func (m Money) Equals(money interface{}) bool {
	switch anyMoney := money.(type) {
	case Dollar:
		return anyMoney.amount == m.amount
	case Franc:
		return anyMoney.amount == m.amount
	default:
		return false
	}
}
