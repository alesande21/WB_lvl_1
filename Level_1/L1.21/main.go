package main

import (
	"context"
	"database/sql"
	"fmt"
)

type DBRepository interface {
	Query(ctx context.Context, query string, args ...any)
	QueryRow(ctx context.Context, query string, args ...any)
}

type sqlLiteDBRepository struct {
	сonn *sql.DB
}

func CreateSqlLiteRepository(newConnection *sql.DB) (DBRepository, error) {
	var rep DBRepository = &sqlLiteDBRepository{сonn: newConnection}
	return rep, nil
}

func (p *sqlLiteDBRepository) Query(ctx context.Context, query string, args ...any) {
	fmt.Println("sqlLiteDBRepository Query")
}

func (p *sqlLiteDBRepository) QueryRow(ctx context.Context, query string, args ...any) {
	fmt.Println("sqlLiteDBRepository QueryRow")
}

type postgresDBRepository struct {
	сonn *sql.DB
}

func CreatePostgresRepository(newConnection *sql.DB) (DBRepository, error) {
	var rep DBRepository = &postgresDBRepository{сonn: newConnection}
	return rep, nil
}

func (p *postgresDBRepository) Query(ctx context.Context, query string, args ...any) {
	fmt.Println("postgresDBRepository Query")
}

func (p *postgresDBRepository) QueryRow(ctx context.Context, query string, args ...any) {
	fmt.Println("postgresDBRepository QueryRow")
}

type RepositoryAdapter struct {
	rep DBRepository
}

func NewRepositoryAdapter(dbRepository DBRepository) *RepositoryAdapter {
	return &RepositoryAdapter{dbRepository}
}

func (ra *RepositoryAdapter) Query(ctx context.Context, query string, args ...any) {
	ra.rep.Query(ctx, query, args...)
}

func (ra *RepositoryAdapter) QueryRow(ctx context.Context, query string, args ...any) {
	ra.rep.Query(ctx, query, args...)
}

func main() {
	var conn sql.DB
	sqlLine, _ := CreateSqlLiteRepository(&conn)
	postgre, _ := CreatePostgresRepository(&conn)

	adapter := NewRepositoryAdapter(sqlLine)
	adapter.QueryRow(context.Background(), "")
	adapter.Query(context.Background(), "")
	adapter = NewRepositoryAdapter(postgre)
	adapter.QueryRow(context.Background(), "")
	adapter.Query(context.Background(), "")

}
