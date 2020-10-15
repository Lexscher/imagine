package models

import (
	"time"
	"context"

	"github.com/lexscher/imagine/pkg/application"
)

type Iuter struct {
	ID        int       `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"-"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
}

func (i *Iuter) Get(ctx context.Context, app *application.Application) error {
	stmt := `SELECT * FROM iuter;`

	err := app.DB.Client.QueryRowContext(
		ctx, 
		stmt, 
	).Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (i *Iuter) GetByID(ctx context.Context, app *application.Application) error {
	stmt := `SELECT * FROM iuter WHERE id = $1;`
	
	err := app.DB.Client.QueryRowContext(
		ctx, 
		stmt, 
		i.ID,
	).Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (i *Iuter) Create(ctx context.Context, app *application.Application) error {
	stmt := `
			INSERT INTO iuter (
				username, 
				password, 
				email
			)
			VALUES ($1, $2, $3)
			RETURNING id;
		`

	err := app.DB.Client.QueryRowContext(
		ctx,
		stmt,
		i.Username,
		i.Password,
		i.Email,
	).Scan(&i.ID)

	if err != nil {
		return err
	}

	return nil
}