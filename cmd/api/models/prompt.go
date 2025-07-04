package models

import (
	"context"
	"deckly/pkg/application"

	"github.com/google/uuid"
)

type Prompt struct {
	ID          uuid.UUID `json:"id"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
}

func (p *Prompt) Create(ctx context.Context, app *application.Application) error {
	stmt := `INSERT INTO prompts (subject, description)VALUES ($1, $2) RETURNING id`
	err := app.DB.Client.QueryRowContext(ctx, stmt, p.Subject, p.Description).Scan(&p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Prompt) GetByID(ctx context.Context, app *application.Application) error {
	stmt := `SELECT * FROM prompts WHERE id = $1`
	err := app.DB.Client.QueryRowContext(ctx, stmt, p.ID).Scan(&p.ID, &p.Subject, &p.Description)
	if err != nil {
		return err
	}
	return nil
}
