package infrastructure

import (
	"errors"

	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/domain/repository"
	"github.com/mattn/go-gimei"
)

type CharacterRepository struct {
}

func NewCharacterRepository() repository.CharacterRepository {
	return &CharacterRepository{}
}

func (cr *CharacterRepository) Create(character *model.Character) (*model.Character, error) {
	var name *gimei.Name
	if character.GetGender() == "" {
		name = gimei.NewName()
	} else if character.GetGender() == "male" {
		name = gimei.NewMale()
	} else if character.GetGender() == "female" {
		name = gimei.NewFemale()
	} else {
		return nil, errors.New("性別に想定外の文字列が設定されています。male・femaleか、何も指定しないでください")
	}
	character, err := model.NewCharacter(name.Sex.String(), name.Last.Kanji(), name.Last.Hiragana(), name.First.Kanji(), name.First.Hiragana())
	if err != nil {
		return nil, err
	}
	return character, nil
}
