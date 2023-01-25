package usecase

import (
	"errors"
	"strings"

	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/domain/repository"
	"golang.org/x/text/width"
)

type DiceUsecase interface {
	Roll(query string) (*model.Dice, error)
}

type diceUsecase struct {
	diceRepository repository.DiceRepository
}

func NewDiceUsecase(diceRepository repository.DiceRepository) DiceUsecase {
	return &diceUsecase{
		diceRepository: diceRepository,
	}
}

func (di diceUsecase) Roll(query string) (*model.Dice, error) {
	// 正規化
	normalizeQuery := strings.ToLower(query)
	normalizeQuery = width.Narrow.String(normalizeQuery)
	// バリデーション
	slice := strings.Split(normalizeQuery, "")
	len := len(slice)
	for i := 0; i < len; i++ {
		if !strings.ContainsAny(slice[i], "0123456789d+") {
			return nil, errors.New("想定外の文字がダイスに指定されました")
		}
	}
	diceModel, err := model.NewDice(normalizeQuery, "", 0)
	if err != nil {
		return nil, err
	}
	dice, err := di.diceRepository.Roll(diceModel)
	if err != nil {
		return nil, err
	}
	return dice, nil
}
