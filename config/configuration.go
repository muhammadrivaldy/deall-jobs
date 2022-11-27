package config

// Configuration is a object configuration
type Configuration struct {
	Port     int    `json:"port" env:"PORT"`
	JWTKey   string `json:"jwt-key" env:"JWT_KEY"`
	Database struct {
		Address    string `json:"address" env:"DB_ADDRESS"`
		User       string `json:"user" env:"DB_USER"`
		Password   string `json:"password" env:"DB_PASSWORD"`
		Parameters string `json:"parameters" env:"DB_PARAMETERS"`
		Schema     struct {
			Security string `json:"security" env:"DB_SCHEMA_SECURITY"`
		} `json:"schema"`
	} `json:"database"`
	EmailSystem struct {
		Host     string `json:"host" env:"EMAIL_HOST"`
		Email    string `json:"email" env:"EMAIL_ADDRESS"`
		Password string `json:"password" env:"EMAIL_PASSWORD"`
		Port     int    `json:"port" env:"EMAIL_PORT"`
	} `json:"email_system"`
	ThirdParty struct {
		Telegram struct {
			Token  string `json:"token" env:"TELEGRAM_TOKEN"`
			ChatID int64  `json:"chat_id" env:"TELEGRAM_CHAT_ID"`
		} `json:"telegram"`
	} `json:"third_party"`
}
