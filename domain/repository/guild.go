package repository

import "github.com/FEATO-org/support-feato-system/domain/model"

type GuildRepository interface {
	Create(guild *model.Guild) (*model.Guild, error)
	FindByID(id int64) (*model.Guild, error)
	Update(guild *model.Guild) (*model.Guild, error)
	Delete(guild *model.Guild) error
}
