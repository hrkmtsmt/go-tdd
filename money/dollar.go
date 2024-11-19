package money

type Dollar struct {
	Money
}

func NewDollar(amount int) Dollar {
	return Dollar{Money: NewMoney(amount)}
}

func (d Dollar) Times(multipler int) Dollar {
	return Dollar{Money: NewMoney(d.amount * multipler)}
}
