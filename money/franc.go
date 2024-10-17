package money

type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	return Franc{amount: amount}
}

func (f *Franc) Times(multipler int) Franc {
	return Franc{
		amount: f.amount * multipler,
	}
}

func (f *Franc) Equals(franc Franc) bool {
	return f.amount == franc.amount
}
