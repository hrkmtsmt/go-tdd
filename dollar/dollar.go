package dollar

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(multipler int) Dollar {
	return Dollar{
		Amount: d.Amount * multipler,
	}
}
