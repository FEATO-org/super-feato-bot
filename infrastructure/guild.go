package infrastructure

import (
	"context"
	"database/sql"

	"github.com/FEATO-org/support-feato-system/db/sqlc"
	"github.com/FEATO-org/support-feato-system/domain/model"
	"github.com/FEATO-org/support-feato-system/domain/repository"
)

type GuildRepository struct {
	db  *sql.DB
	ctx context.Context
}

// Create implements repository.GuildRepository.
func (gr *GuildRepository) Create(guild *model.Guild) (*model.Guild, error) {
	queries := sqlc.New(gr.db)

	result, err := queries.CreateGuild(gr.ctx, sqlc.CreateGuildParams{
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
	if err = guild.Set(id, guild.Name, guild.DiscordID, guild.SheetID); err != nil {
		return nil, err
	}
	return guild, nil
}

// Delete implements repository.GuildRepository.
func (*GuildRepository) Delete(guild *model.Guild) error {
	panic("unimplemented")
}

// FindByID implements repository.GuildRepository.
func (*GuildRepository) FindByID(id int64) (*model.Guild, error) {
	panic("unimplemented")
}

// Update implements repository.GuildRepository.
func (*GuildRepository) Update(guild *model.Guild) (*model.Guild, error) {
	panic("unimplemented")
}

func NewGuildRepository(db *sql.DB, ctx context.Context) repository.GuildRepository {
	return &GuildRepository{
		db:  db,
		ctx: ctx,
	}
}
