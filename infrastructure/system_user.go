package infrastructure

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FEATO-org/support-feato-system/db/sqlc"
	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/domain/repository"
)

type SystemUserRepository struct {
	db  *sql.DB
	ctx context.Context
}

// Create implements repository.SystemUserRepository.
func (su *SystemUserRepository) Create(systemUser *model.SystemUser) (*model.SystemUser, error) {
	queries := sqlc.New(su.db)

	result, err := queries.CreateSystemUser(su.ctx, systemUser.DiscordID)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	if err = systemUser.Set(id, systemUser.DiscordID, systemUser.Guilds); err != nil {
		return nil, err
	}

	for _, guild := range systemUser.Guilds {
		result, err := queries.FindByDiscordIDGuild(su.ctx, guild.DiscordID)
		if errors.Is(err, sql.ErrNoRows) {
			result, err := queries.CreateGuild(su.ctx, sqlc.CreateGuildParams{
				Name:      guild.Name,
				DiscordID: guild.DiscordID,
				SheetID: sql.NullString{
					String: guild.SheetID,
				},
			})
			if err != nil {
				return nil, err
			}

			id, err := result.LastInsertId()
			if err != nil {
				return nil, err
			}
			err = guild.Set(id, guild.Name, guild.DiscordID, guild.SheetID)
			if err != nil {
				return nil, err
			}
		}
		if err != nil {
			return nil, err
		}

		err = guild.Set(result.ID, result.Name, guild.DiscordID, guild.SheetID)
		if err != nil {
			return nil, err
		}

		_, err = queries.CreateSystemUserGuild(su.ctx, sqlc.CreateSystemUserGuildParams{
			SystemUserID: systemUser.Id,
			GuildID:      guild.Id,
		})
		if err != nil {
			return nil, err
		}
	}

	return systemUser, nil
}

// Delete implements repository.SystemUserRepository.
func (*SystemUserRepository) Delete(systemUser *model.SystemUser) error {
	panic("unimplemented")
}

// FindByID implements repository.SystemUserRepository.
func (*SystemUserRepository) FindByID(id int64) (*model.SystemUser, error) {
	panic("unimplemented")
}

// Update implements repository.SystemUserRepository.
func (*SystemUserRepository) Update(systemUser *model.SystemUser) (*model.SystemUser, error) {
	panic("unimplemented")
}

func NewSystemUserRepository(db *sql.DB, ctx context.Context) repository.SystemUserRepository {
	return &SystemUserRepository{
		db:  db,
		ctx: ctx,
	}
}
