package main

import (
	"context"
	"database/sql"
	"fmt"
)

// DBRepository Определяем интерфейс DBRepository, который содержит методы Query и QueryRow
// Этот интерфейс представляет универсальный интерфейс для работы с различными базами данных
type DBRepository interface {
	Query(ctx context.Context, query string, args ...any)
	QueryRow(ctx context.Context, query string, args ...any)
}

// Структура sqlLiteDBRepository представляет реализацию интерфейса для базы данных SQLite
type sqlLiteDBRepository struct {
	сonn *sql.DB
}

// CreateSqlLiteRepository Функция CreateSqlLiteRepository создает новый экземпляр sqlLiteDBRepository
// и возвращает его как DBRepository
func CreateSqlLiteRepository(newConnection *sql.DB) (DBRepository, error) {
	var rep DBRepository = &sqlLiteDBRepository{сonn: newConnection}
	return rep, nil
}

// Query Метод Query иммитирует выполнение запроса
func (p *sqlLiteDBRepository) Query(ctx context.Context, query string, args ...any) {
	fmt.Println("sqlLiteDBRepository Query")
}

// QueryRow Метод QueryRow иммитирует выполнение запроса
func (p *sqlLiteDBRepository) QueryRow(ctx context.Context, query string, args ...any) {
	fmt.Println("sqlLiteDBRepository QueryRow")
}

// Структура postgresDBRepository представляет реализацию интерфейса для базы данных postgres
type postgresDBRepository struct {
	сonn *sql.DB
}

// CreatePostgresRepository Функция CreatePostgresRepository создает новый экземпляр postgresDBRepository
// и возвращает его как DBRepository
func CreatePostgresRepository(newConnection *sql.DB) (DBRepository, error) {
	var rep DBRepository = &postgresDBRepository{сonn: newConnection}
	return rep, nil
}

// Query Метод Query иммитирует выполнение запроса
func (p *postgresDBRepository) Query(ctx context.Context, query string, args ...any) {
	fmt.Println("postgresDBRepository Query")
}

// QueryRow Метод QueryRow иммитирует выполнение запроса
func (p *postgresDBRepository) QueryRow(ctx context.Context, query string, args ...any) {
	fmt.Println("postgresDBRepository QueryRow")
}

// RepositoryAdapter Структура RepositoryAdapter представляет собой адаптер,
// который использует универсальный интерфейс DBRepository для выполнения операций с базами данных
type RepositoryAdapter struct {
	rep DBRepository
}

// NewRepositoryAdapter создает новый экземпляр адаптера
func NewRepositoryAdapter(dbRepository DBRepository) *RepositoryAdapter {
	return &RepositoryAdapter{dbRepository}
}

// Query вызывает Query у конкретной реализации DBRepository
func (ra *RepositoryAdapter) Query(ctx context.Context, query string, args ...any) {
	ra.rep.Query(ctx, query, args...)
}

// QueryRow вызывает QueryRow у конкретной реализации DBRepository
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
