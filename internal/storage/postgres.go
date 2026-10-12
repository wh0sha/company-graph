package storage

import (
	"company-graph/internal/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage(ctx context.Context, connString string) (*Storage, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать пул полдключения: %w", err)
	}
	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к бд: %w", err)
	}

	return &Storage{pool: pool}, nil
}

func (s *Storage) SaveCompanies(ctx context.Context, companies []model.Company) error {
	query := `INSERT INTO companies (name, inn, ogrn, address, source) VALUES ($1, $2, $3, $4, $5)`

	for _, comp := range companies {
		_, err := s.pool.Exec(ctx, query, comp.Name, comp.INN, comp.OGRN, comp.Address, comp.Source)
		if err != nil {
			return fmt.Errorf("ошибка сохранения компании %s: %w", comp.Name, err)
		}
	}
	return nil
}

func (s *Storage) GetCompanies(ctx context.Context) ([]model.Company, error) {
	query := `SELECT id, name, inn, ogrn, address, source, created_at FROM companies`

	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к БД: %w", err)
	}
	defer rows.Close()

	var companies []model.Company

	for rows.Next() {
		var comp model.Company

		err = rows.Scan(&comp.ID, &comp.Name, &comp.INN, &comp.OGRN, &comp.Address, &comp.Source, &comp.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения строки: %w", err)
		}
		companies = append(companies, comp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка после чтения строк: %w", err)
	}

	return companies, nil
}
