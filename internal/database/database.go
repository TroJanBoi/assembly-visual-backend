package database

import (
	"fmt"
	"log"
	"os"

	"github.com/TroJanBoi/assembly-visual-backend/internal/conf"
	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service interface {
	GetClient() *gorm.DB

	Close() error
}

type service struct {
	db *gorm.DB
}

func (s *service) GetClient() *gorm.DB {
	return s.db
}

var dbInstance *service

func New() Service {
	config := conf.NewConfig()
	if dbInstance != nil {
		return dbInstance
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Error,
			Colorful: true,
		},
	)
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
		log.Fatalf("Error connecting to database: %v", err)
		return nil
	}
	if config.AUTO_MIGRATE {
		log.Println("Auto migrating database...")
		err = db.AutoMigrate(
			// Add your models here, e.g. &model.User{}, etc.
			&model.User{},
		)
		if err != nil {
			log.Fatalf("Error auto migrating database: %v", err)
			return nil
		}
		log.Println("Database auto migration completed successfully.")
	} else {
		log.Println("Auto migration is disabled. Skipping database migration.")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting sql.DB from gorm.DB: %v", err)
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetConnMaxLifetime(10)
	sqlDB.SetMaxOpenConns(10)
	err = sqlDB.Ping()
	if err != nil {
		log.Panic("Could not ping database")
		return nil
	}
	dbInstance = &service{db: db}
	return dbInstance
}

func (s *service) Close() error {
	config := conf.NewConfig()
	log.Printf("Closing database connection to %s", config.POSTGRES_DB)
	sqlDB, err := s.db.DB()
	if err != nil {
		log.Printf("Error getting sql.DB from gorm.DB: %v", err)
		return err
	}
	return sqlDB.Close()
}
