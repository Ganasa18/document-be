package appconfig

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	UserModel "github.com/Ganasa18/document-be/internal/auth/model/domain"
	RoleModel "github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *Config, config *gorm.Config) (*gorm.DB, error) {
	tz := "Asia/Jakarta"

	if config == nil {
		config = &gorm.Config{}
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password='%s' dbname=%s sslmode=disable TimeZone=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUsername,
		cfg.DbPass,
		cfg.DbName,
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

	db.AutoMigrate(&UserModel.UserModel{}, &RoleModel.RoleMasterModel{})

	return db, nil

	// migrate create -ext sql -dir scripts/migrations create_table_first
	// migrate -database "postgres://postgres:admin@localhost:5432/db_document_builder?sslmode=disable" -path scripts/migrations up
}
