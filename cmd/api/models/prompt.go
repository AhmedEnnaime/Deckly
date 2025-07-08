package models

import (
	"bytes"
	"context"
	"deckly/pkg/application"
	"encoding/json"
	"net/http"
	"time"

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

func (p *Prompt) TriggerN8nWorkflow(app *application.Application) error {
	url := app.Cfg.GetN8NWebhookURL()
	client := &http.Client{Timeout: 5 * time.Second}
	body, err := json.Marshal(p)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	return err
}
