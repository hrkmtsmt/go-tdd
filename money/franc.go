package money

type Franc struct {
	Money
}

func NewFranc(amount int) Franc {
	return Franc{
		Money: NewMoney(amount),
	}
}

func (f Franc) Times(multipler int) Franc {
	return Franc{
		Money: NewMoney(f.amount * multipler),
	}
}
