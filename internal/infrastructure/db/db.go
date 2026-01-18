package db

import (
	"context"
	"fmt"
	"main/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Conn struct {
	*pgxpool.Pool
}

func (c *Conn) GetPool() *pgxpool.Pool {
	return c.Pool
}

func (c *Conn) Close() {
	c.Pool.Close()
}

func Connect() (*Conn, error) {
	dsn := config.GetDbConnectionString().ConnectionString
	if dsn == "" {
		return nil, fmt.Errorf("database connection string is empty")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to create database pool: %v", err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return &Conn{pool}, nil
}
