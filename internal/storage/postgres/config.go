package postgres

// Config ..
type Config struct {
	Env      string `env:"ENVIRONMENT,default=test"`
	User     string `env:"POSTGRES_USER,default=aliouamar"`
	Password string `env:"POSTGRES_PASSWORD,default=v9bnVv31n"`
	Host     string `env:"POSTGRES_HOST,default=localhost"`
	Port     int    `env:"POSTGRES_PORT,default=5432"`
	DbName   string `env:"POSTGRES_DBNAME,default=shilling_dev"`
}

// NewConfig ...
func NewConfig(env, user, password, host, dbName string, port int) *Config {
	return &Config{
		Env:      env,
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		DbName:   dbName,
	}
}
