package model

type Character struct {
	gender            string
	firstName         string
	firstNameHiragana string
	lastName          string
	lastNameHiragana  string
}

func NewCharacter(gender, firstName, firstNameHiragana,
	lastName, lastNameHiragana string) (*Character, error) {
	return &Character{
		gender:            gender,
		firstName:         firstName,
		firstNameHiragana: firstNameHiragana,
		lastName:          lastName,
		lastNameHiragana:  lastNameHiragana,
	}, nil
}

func (c *Character) Set(gender, firstName, firstNameHiragana,
	lastName, lastNameHiragana string) error {
	c.gender = gender
	c.firstName = firstName
	c.firstNameHiragana = firstNameHiragana
	c.lastName = lastName
	c.lastNameHiragana = lastNameHiragana
	return nil
}

func (c *Character) GetGender() string {
	return c.gender
}

func (c *Character) GetName() string {
	return c.firstName + c.lastName
}

func (c *Character) GetNameHiragana() string {
	return c.firstNameHiragana + c.lastNameHiragana
}
