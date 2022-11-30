package config

// Configuration is a object configuration
type Configuration struct {
	Port     int    `json:"port" env:"PORT"`
	JWTKey   string `json:"jwt_key" env:"JWT_KEY"`
	Database struct {
		User       string `json:"user" env:"DB_USER"`
		Password   string `json:"password" env:"DB_PASSWORD"`
		Parameters string `json:"parameters" env:"DB_PARAMETERS"`
		Schema     struct {
			Security struct {
				Address       string `json:"address" env:"DB_ADDRESS_SECURITY"`
				Database      string `json:"database" env:"DB_SCHEMA_SECURITY"`
				MigrationPath string `json:"migration_path" env:"DB_MIGRATION_PATH_SECURITY"`
			} `json:"security"`
			User struct {
				Address       string `json:"address" env:"DB_ADDRESS_USER"`
				Database      string `json:"database" env:"DB_SCHEMA_USER"`
				MigrationPath string `json:"migration_path" env:"DB_MIGRATION_PATH_USER"`
			} `json:"user"`
		} `json:"schema"`
	} `json:"database"`
	Redis struct {
		Address  string `json:"address" env:"REDIS_ADDRESS"`
		Password string `json:"password" env:"REDIS_PASSWORD"`
	} `json:"redis"`
	ThirdParty struct {
		Telegram struct {
			Token  string `json:"token" env:"TELEGRAM_TOKEN"`
			ChatID int64  `json:"chat_id" env:"TELEGRAM_CHAT_ID"`
		} `json:"telegram"`
	} `json:"third_party"`
}
