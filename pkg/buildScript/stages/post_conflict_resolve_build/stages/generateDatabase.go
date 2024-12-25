package post_conflict_resolve_stages

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	"github.com/charmbracelet/log"
)

func writeDatabaseSchemaFile() {
	log.Warn("This writeDatabaseSchemaFile method is only here for future-proofing. It currently does nothing.")
}

// Writes the schema file and then runs prisma related scripts and prepares the database to be seeded.
func GenerateDatabase(cfg *build_config.BuildManager) {
	writeDatabaseSchemaFile()
	cfg.Db.GenerateFromStrategy(cfg.PackageManager())
}
