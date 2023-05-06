package repository

import (
	"github.com/FEATO-org/support-feato-system/domain/model"
)

type DiceRepository interface {
	Roll(dice *model.Dice) (*model.Dice, error)
}
