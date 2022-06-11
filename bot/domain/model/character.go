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

// 日本人名用の組み立て
func (c *Character) GetName() string {
	return c.lastName + " " + c.firstName
}

func (c *Character) GetNameHiragana() string {
	return c.lastNameHiragana + " " + c.firstNameHiragana
}

func (c *Character) GetGender() string {
	return c.gender
}
