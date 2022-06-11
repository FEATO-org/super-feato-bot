package model

import "strconv"

type Dice struct {
	query   string
	formula string
	total   int
}

func NewDice(query, formula string, total int) (*Dice, error) {
	return &Dice{
		query:   query,
		formula: formula,
		total:   total,
	}, nil
}

func (d *Dice) Set(query, formula string, total int) error {
	d.query = query
	d.formula = formula
	d.total = total
	return nil
}

func (d *Dice) GetResult() string {
	return d.query + "：" + strconv.Itoa(d.total) + " ( " + d.formula + " )"
}

func (d *Dice) GetQuery() string {
	return d.query
}
