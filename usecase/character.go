package usecase

import (
	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/domain/repository"
)

type CharacterUsecase interface {
	Generate(gender string) (*model.Character, error)
}

type characterUsecase struct {
	characterRepository repository.CharacterRepository
}

// Generate implements CharacterUsecase
func (cu *characterUsecase) Generate(gender string) (*model.Character, error) {
	character, err := model.NewCharacter(gender, "", "", "", "")
	if err != nil {
		return nil, err
	}
	return cu.characterRepository.Create(character)
}

func NewCharacterUsecase(characterRepository repository.CharacterRepository) CharacterUsecase {
	return &characterUsecase{
		characterRepository: characterRepository,
	}
}
