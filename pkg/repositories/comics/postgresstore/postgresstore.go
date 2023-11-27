package postgresstore

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/caarlos0/env/v10"
	_ "github.com/lib/pq"

	datalayer "github.com/klaital/klaital.com/pkg/repositories/comics"
	"github.com/klaital/klaital.com/pkg/repositories/comics/postgresstore/queries"
	"github.com/klaital/klaital.com/pkg/repositories/utilities"
)

//go:embed migrations/*
var MigrationsFS embed.FS

type Repository struct {
	Queries *queries.Queries
}

type Config struct {
	DbName              string `env:"COMICS_DB_NAME" envDefault:"webcomics"`
	ConnectionStringFmt string `env:"COMICS_DB_CONNECT_FMT" envDefault:"postgres://klaital:rootpw@postgres-klaital:5432/%s?sslmode=disable"`
}

func New() (*Repository, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parsing comics_postgresstore env: %w", err)
	}

	db, err := utilities.ConnectAndMigratePostgres(cfg.ConnectionStringFmt, cfg.DbName, MigrationsFS)
	if err != nil {
		return nil, fmt.Errorf("connecting to db for comics_postgresstore: %w", err)
	}

	// Initialize Repository
	repo := &Repository{
		Queries: queries.New(db),
	}

	return repo, nil
}

func StringPtrToSql(s *string) sql.NullString {
	ns := sql.NullString{
		String: "",
		Valid:  false,
	}
	if s != nil {
		ns.Valid = true
		ns.String = *s
	}
	return ns
}

func (s *Repository) AddComic(ctx context.Context, c *datalayer.Comic) error {
	params := queries.AddComicParams{
		Ordinal: sql.NullInt32{
			Valid: true,
			Int32: int32(c.Ordinal),
		},
		UserID:           int32(c.UserID),
		Title:            c.Title,
		BaseUrl:          sql.NullString{Valid: true, String: c.BaseURL},
		FirstComicUrl:    StringPtrToSql(c.FirstComicUrl),
		LatestComicUrl:   StringPtrToSql(c.LatestComicUrl),
		RssUrl:           StringPtrToSql(c.RssUrl),
		UpdatesMonday:    c.UpdatesMonday,
		UpdatesTuesday:   c.UpdatesTuesday,
		UpdatesWednesday: c.UpdatesWednesday,
		UpdatesThursday:  c.UpdatesThursday,
		UpdatesFriday:    c.UpdatesFriday,
		UpdatesSaturday:  c.UpdatesSaturday,
		UpdatesSunday:    c.UpdatesSunday,
		Active:           c.Active,
		Nsfw:             c.Nsfw,
		LastRead:         sql.NullTime{Valid: false},
	}

	err := s.Queries.AddComic(ctx, params)
	if err != nil {
		return fmt.Errorf("adding comic: %w", err)
	}

	//Success!
	return nil
}

func (s *Repository) GetComics(ctx context.Context, userId uint64, filterNsfw *bool, filterActive *bool) ([]datalayer.Comic, error) {
	return nil, datalayer.ErrNotImplemented
}

func (s *Repository) GetComic(ctx context.Context, userId, comicId uint64) (*datalayer.Comic, error) {
	return nil, datalayer.ErrNotImplemented
}

func (s *Repository) UpdateComic(ctx context.Context, c *datalayer.Comic) error {
	return datalayer.ErrNotImplemented
}

func (s *Repository) DeleteComic(ctx context.Context, userId, comicId uint64) error {
	return datalayer.ErrNotImplemented
}

func (s *Repository) MarkComicRead(ctx context.Context, userId, comicId uint64) error {
	return datalayer.ErrNotImplemented
}

func (s *Repository) AddRssItems(ctx context.Context, userId, comicId uint64, rssItems []datalayer.RssItem) error {
	return datalayer.ErrNotImplemented
}

func (s *Repository) MarkRssRead(ctx context.Context, userId uint64, rssItemGuid string) error {
	return datalayer.ErrNotImplemented
}
