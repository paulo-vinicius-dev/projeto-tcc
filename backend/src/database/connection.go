package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// Connection ...
func Connection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err := conn.Ping(context.Background()); err != nil {
		conn.Close(context.Background())
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
