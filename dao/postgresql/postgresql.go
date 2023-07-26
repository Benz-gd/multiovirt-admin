package postgresql

import (
	"fmt"
	"go.uber.org/zap"
	postgresql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"multiovirt-admin/settings"
	"strconv"
	"time"
)

var PG *gorm.DB

func InitPostgreSQL(cfg *settings.PostgreSQLConfig) (*gorm.DB, error) {

	pgprestatement, err := strconv.ParseBool(cfg.PGPreStatement)
	if err != nil {
		zap.L().Error("InitPostgreSQL PGPreStatement error!", zap.Error(err))
	}
	dsn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.TimeZone)
	db, err := gorm.Open(postgresql.New(postgresql.Config{DSN: dsn, PreferSimpleProtocol: pgprestatement}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.PGMaxIdelConns)
	sqlDB.SetMaxOpenConns(cfg.PGMaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.PGConnMaxLifetime) * time.Minute)
	return db, nil
}

func DBClose() {
	sqlDB, err := PG.DB()
	if err != nil {
		zap.L().Error("func DBClose: ", zap.Error(err))
	}
	sqlDB.Close()
}
