package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"log"
	"os"
)

type Config struct {
	HTTP  HTTPConf
	DB    PostgresConfig
	Excel ExcelServerConf

	CtxTimeout int

	LogLevel string
}

type HTTPConf struct {
	Host string
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type ExcelServerConf struct {
	Host string
	Port string
	URL  string
}

func Load(envFile string) (cfg Config) {

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal(err)
		return
	}

	cfg.HTTP.Host = cast.ToString(os.Getenv("HTTP_HOST"))
	cfg.HTTP.Port = cast.ToString(os.Getenv("HTTP_PORT"))

	cfg.DB.Host = cast.ToString(os.Getenv("POSTGRES_HOST"))
	cfg.DB.Port = cast.ToString(os.Getenv("POSTGRES_PORT"))
	cfg.DB.Username = cast.ToString(os.Getenv("POSTGRES_USERNAME"))
	cfg.DB.Password = cast.ToString(os.Getenv("POSTGRES_PASSWORD"))
	cfg.DB.DBName = cast.ToString(os.Getenv("POSTGRES_DB"))
	cfg.DB.SSLMode = cast.ToString(os.Getenv("POSTGRES_SSLMODE"))

	cfg.Excel.Host = cast.ToString(os.Getenv("EXCEL_HOST"))
	cfg.Excel.Port = cast.ToString(os.Getenv("EXCEL_PORT"))
	cfg.Excel.URL = fmt.Sprintf("http://%s:%s/common/getform", cfg.Excel.Host, cfg.Excel.Port)

	cfg.CtxTimeout = cast.ToInt(os.Getenv("CTX_TIMEOUT"))

	cfg.LogLevel = cast.ToString(os.Getenv("LOG_LEVEL"))

	return cfg
}
