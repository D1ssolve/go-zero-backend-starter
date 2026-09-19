package note

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) Create(ctx context.Context, text string) (Note, error) {
	const query = `
		INSERT INTO notes (text)
		VALUES ($1)
		RETURNING id, text, created_at`

	var result Note
	err := r.pool.QueryRow(ctx, query, text).Scan(&result.ID, &result.Text, &result.CreatedAt)
	return result, err
}

func (r *PostgresRepository) List(ctx context.Context, limit int) ([]Note, error) {
	const query = `
		SELECT id, text, created_at
		FROM notes
		ORDER BY id DESC
		LIMIT $1`

	rows, err := r.pool.Query(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]Note, 0, limit)
	for rows.Next() {
		var item Note
		if err := rows.Scan(&item.ID, &item.Text, &item.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, rows.Err()
}
