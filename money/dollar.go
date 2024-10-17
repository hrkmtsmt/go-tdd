package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{amount: amount}
}

func (d *Dollar) Times(multipler int) Dollar {
	return Dollar{
		amount: d.amount * multipler,
	}
}

func (d *Dollar) Equals(dollar Dollar) bool {
	return d.amount == dollar.amount
}
