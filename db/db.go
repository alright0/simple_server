package db

import (
	"context"
	"fmt"
	"main/config"

	"github.com/jackc/pgx/v5"
)

type Conn struct {
	*pgx.Conn
}

func (c *Conn) GetPgConn() *pgx.Conn {
	return c.Conn
}

func (c *Conn) Close() {
	c.Conn.Close(context.Background())
}

func Connect() (*Conn, error) {
	conn, err := pgx.Connect(context.Background(), config.GetDbConnectionString().ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	return &Conn{conn}, nil
}
