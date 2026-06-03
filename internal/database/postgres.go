package database

import (
    "context"
    "github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
    return pgx.Connect(
        context.Background(),
        "postgres://postgres:postgres@localhost:5432/user_management",
    )
}