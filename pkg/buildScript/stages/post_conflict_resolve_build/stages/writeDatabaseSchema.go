package post_conflict_resolve_stages

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
)

// Writes the schema file and then runs prisma related scripts and prepares the database to be seeded.
func WriteDatabaseSchemaFile(cfg *build_config.BuildManager) {
	cfg.Db.WriteSchemaFile(cfg.Paths.PrismaSchema())
}
