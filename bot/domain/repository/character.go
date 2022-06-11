package repository

import "github.com/FEATO-org/support-feato-system/domain/model"

type CharacterRepository interface {
	Generate(character *model.Character) (*model.Character, error)
}
