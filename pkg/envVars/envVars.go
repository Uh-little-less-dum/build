package env_vars

type EnvVariable string

// TODO: Replace all calls to os.Getenv with these variables for improved reliability.
const (
	AdditionalSources EnvVariable = "ULLD_ADDITIONAL_SOURCES"
	LogLevel          EnvVariable = "ULLD_LOG_LEVEL"
	LogFile           EnvVariable = "ULLD_LOG_FILE"
)
