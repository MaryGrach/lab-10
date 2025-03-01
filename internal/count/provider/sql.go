package provider

import (
	"database/sql"
	"errors"

	"github.com/Degreezee/labwork10/pkg/vars"
)

func (p *Provider) SelectCount() (int, error) {
	var msg int
	err := p.conn.QueryRow("SELECT number FROM counter WHERE id_number = 1").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, vars.ErrorDBNotInitialized
		}
		return 0, err
	}
	return msg, nil
}

func (p *Provider) UpdateCount(msg int) error {
	_, err := p.conn.Exec("UPDATE counter SET number = number + $1 WHERE id_number = 1", msg)
	if err != nil {
		return err
	}

	return nil
}
