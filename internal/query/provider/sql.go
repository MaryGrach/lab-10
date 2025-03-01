package provider

import (
	"database/sql"
	"errors"

	"github.com/Degreezee/labwork10/pkg/vars"
)

func (p *Provider) SelectQuery() (string, error) {
	var msg string
	err := p.conn.QueryRow("SELECT name_query FROM query").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", vars.ErrorDBNotInitialized
		}
		return "", err
	}
	return msg, nil
}

func (p *Provider) UpdateQuery(msg string) error {
	_, err := p.conn.Exec("UPDATE query SET name_query = $1", msg)
	if err != nil {
		return err
	}

	return nil
}
