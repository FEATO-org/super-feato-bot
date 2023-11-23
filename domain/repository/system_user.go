package repository

import "github.com/FEATO-org/support-feato-system/domain/model"

type SystemUserRepository interface {
	Create(systemUser *model.SystemUser) (*model.SystemUser, error)
	FindByID(id int64) (*model.SystemUser, error)
	Update(systemUser *model.SystemUser) (*model.SystemUser, error)
	Delete(systemUser *model.SystemUser) error
}
