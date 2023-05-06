package repository

import "github.com/FEATO-org/support-feato-system/domain/model"

type CharacterRepository interface {
	Create(character *model.Character) (*model.Character, error)
}
