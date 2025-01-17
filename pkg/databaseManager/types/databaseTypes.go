package database_types

type DatabaseType string

// Setup sqlite here as well.
// Enum values must match database schema filename.
const (
	Postgres DatabaseType = "postgres"
)
