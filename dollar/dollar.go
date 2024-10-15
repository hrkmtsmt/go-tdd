package dollar

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(multipler int) Dollar {
	return Dollar{
		Amount: d.Amount * multipler,
	}
}

func (d *Dollar) Equals(dollar Dollar) bool {
	return d.Amount == dollar.Amount
}
