package conf

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"path/filepath"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConfigInterface interface {
	CreateClientDatabase() (interface{}, interface{}, error)
}

type Config struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_SSL      string
	POSTGRES_TIMEZONE string
	FE_URL            string
	ENV               string
	PORT              int
	AUTO_MIGRATE      bool
}

func GetGoogleOAuthConfig() *oauth2.Config {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("GOOGLE_REDIRECT_URI")
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"openid",
			"email",
			"profile",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrThrow(key string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func getEnvBoolOrDefault(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	switch value {
	case "true", "TRUE", "1", "yes", "YES", "on", "ON":
		return true
	case "false", "FALSE", "0", "no", "NO", "off", "OFF":
		return false
	default:
		log.Fatalf("Invalid value for boolean environment variable %s: %s", key, value)
		return defaultValue
	}
}

func (c *Config) CreateClientDatabase() (*gorm.DB, *sql.DB, error) {
	config := NewConfig()
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel: logger.Info,
		Colorful: true,
	})

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.POSTGRES_HOST,
		config.POSTGRES_PORT,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_DB,
		config.POSTGRES_SSL,
		config.POSTGRES_TIMEZONE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panicf("Error connecting to database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Panicf("Error getting sql.DB from gorm.DB: %v", err)
	}
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetConnMaxLifetime(10)
	sqlDB.SetMaxOpenConns(100)

	err = sqlDB.Ping()
	if err != nil {
		log.Panicf("Error pinging database: %v", err)
	}

	return db, sqlDB, nil
}

func NewConfig() *Config {
	// dotenverr := godotenv.Load()
	// if dotenverr != nil {
	// 	log.Printf("Warning: .env file not found, using environment variables instead: %v", dotenverr)
	// } else {
	// 	log.Println("Loaded environment variables from .env file")
	// }
	paths := []string{
		".env",
		filepath.Join("..", "..", ".env"),
	}
	loaded := false
	for _, p := range paths {
		if err := godotenv.Load(p); err == nil {
			log.Printf("Loaded environment variables from %s", p)
			loaded = true
			break
		}
	}
	if !loaded {
		log.Printf("Warning: .env file not found in %v, using OS environment", paths)
	}

	portStr := getEnvOrDefault("PORT", "9090")
	portInt := 9090
	if p, err := strconv.Atoi(portStr); err == nil {
		portInt = p
	} else {
		log.Printf("Warning: Invalid PORT value '%s', using default 9090", portStr)
	}

	return &Config{
		POSTGRES_USER:     getEnvOrThrow("POSTGRES_USER"),
		POSTGRES_PASSWORD: getEnvOrThrow("POSTGRES_PASSWORD"),
		POSTGRES_DB:       getEnvOrThrow("POSTGRES_DB"),
		POSTGRES_HOST:     getEnvOrDefault("POSTGRES_HOST", "postgres"),
		POSTGRES_PORT:     getEnvOrDefault("POSTGRES_PORT", "5432"),
		POSTGRES_SSL:      getEnvOrDefault("POSTGRES_SSL", "disable"),
		POSTGRES_TIMEZONE: getEnvOrDefault("POSTGRES_TIMEZONE", "Asia/Bangkok"),
		ENV:               getEnvOrDefault("ENV", "dev"),
		FE_URL:            getEnvOrDefault("FE_URL", "http://localhost:3000"),
		PORT:              portInt,
		AUTO_MIGRATE:      getEnvBoolOrDefault("AUTO_MIGRATE", false),
	}

}
