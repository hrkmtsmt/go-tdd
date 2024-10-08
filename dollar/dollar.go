package dollar

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(multipler int) {
	d.Amount *= multipler
}
