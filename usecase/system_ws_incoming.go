package usecase

import (
	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/domain/repository"
)

type SystemWSIncomingUsecase interface {
	ReceiveEEW(body interface{}, test bool) (*model.EEW, error)
}

type systemWSIncomingUsecase struct {
	eewRepository repository.EEWRepository
}

func NewSystemWSIncoming(eewRepository repository.EEWRepository) SystemWSIncomingUsecase {
	return &systemWSIncomingUsecase{
		eewRepository: eewRepository,
	}
}

func (su systemWSIncomingUsecase) ReceiveEEW(body interface{}, test bool) (*model.EEW, error) {
	return model.NewEEW(body, false)
}
