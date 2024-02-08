package appconfig

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Ganasa18/document-be/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *Config, config *gorm.Config) (*gorm.DB, error) {
	tz := "Asia/Jakarta"

	if config == nil {
		config = &gorm.Config{}
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password='%s' dbname=%s search_path=%s sslmode=disable TimeZone=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUsername,
		cfg.DbPass,
		cfg.DbName,
		cfg.DbSchema,
		tz,
	)
	sqlDB, err := sql.Open("pgx", dsn)
	utils.IsErrorDoPanic(err)

	maxConn := os.Getenv(utils.CONFIG_DB_MAX_CONNECTION)
	c, err := strconv.Atoi(maxConn)
	utils.IsErrorDoPanic(err)

	sqlDB.SetMaxOpenConns(c)

	maxIdle := os.Getenv(utils.CONFIG_DB_MAX_IDLE_CONNECTION)
	i, err := strconv.Atoi(maxIdle)
	utils.IsErrorDoPanic(err)

	sqlDB.SetMaxIdleConns(i)

	lifeTime := os.Getenv(utils.CONFIG_DB_MAX_LIFETIME_CONNECTION)
	lt, err := time.ParseDuration(lifeTime)
	utils.IsErrorDoPanic(err)

	sqlDB.SetConnMaxLifetime(lt)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.IsErrorDoPanic(err)
	}

	return db, nil
}
