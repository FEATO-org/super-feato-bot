package model

type EEW struct {
	source interface{}
	test   bool
}

func NewEEW(source interface{}, test bool) (*EEW, error) {
	return &EEW{
		source: source,
		test:   test,
	}, nil
}

func (e *EEW) Set(source interface{}, test bool) error {
	e.source = source
	e.test = test
	return nil
}
