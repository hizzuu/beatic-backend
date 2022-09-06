package repository

import (
	"context"
	"database/sql"

	"github.com/hizzuu/beatic-backend/internal/domain"
)

type userRepo struct {
	db db
}

func NewUser(db db) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Get(ctx context.Context, id int) (*domain.User, error) {
	q := "SELECT id, name, gender, birthday FROM users WHERE id = ?"

	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	user, err := convertRowToUser(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	query := `INSERT users SET name=?, gender=?, birthday=?`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, u.Name, u.Gender, u.Birthday)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = id

	return u, nil
}

func convertRowToUser(row *sql.Row) (*domain.User, error) {
	user := &domain.User{}
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Gender,
		&user.Birthday,
	); err != nil {
		return nil, err
	}

	return user, nil
}
