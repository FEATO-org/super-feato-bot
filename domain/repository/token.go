package repository

import "github.com/FEATO-org/support-feato-system/domain/model"

type TokenRepository interface {
	Create(token *model.Token) (*model.Token, error)
	FindByID(id int64) (*model.Token, error)
	Update(token *model.Token) (*model.Token, error)
	Delete(token *model.Token) error
}
