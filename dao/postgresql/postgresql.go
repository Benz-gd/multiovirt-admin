package postgresql

import (
	"fmt"
	"example/fundemo01/web-app/settings"
	"go.uber.org/zap"
	postgresql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var PG *gorm.DB



func  InitPostgreSQL(cfg *settings.PostgreSQLConfig) (*gorm.DB,error) {
	//self.pgaddr = config.Conf.PostgreSQL.PGAddr
	//self.pgport = config.Conf.PostgreSQL.PGPort
	//self.pguser = config.Conf.PostgreSQL.PGUser
	//self.pgpwd = config.Conf.PostgreSQL.PGPWD
	//self.pgdb = config.Conf.PostgreSQL.PGDB
	//self.pgtz = config.Conf.PostgreSQL.PGTZ
	//self.pgprestatement = config.Conf.PostgreSQL.PGPreStatement
	//self.pgmaxidelconns = config.Conf.PostgreSQL.PGMaxIdleConns
	//self.pgmaxopenconns = config.Conf.PostgreSQL.PGMaxOpenConns
	//self.pgconnmaxlifetime = time.Duration(config.Conf.PostgreSQL.PGConnMaxLifeTime)*time.Minute
	//self.pgconnmaxlifetime = sec.Key("PGConnMaxLifetime").Duration()
	pgprestatement,err := strconv.ParseBool(cfg.PGPreStatement)
	if err != nil{
		zap.L().Error("InitPostgreSQL PGPreStatement error!",zap.Error(err))
	}
	dsn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",cfg.Host,cfg.User,cfg.Password,cfg.DBName,cfg.Port,cfg.TimeZone)
	db,err := gorm.Open(postgresql.New(postgresql.Config{DSN:dsn,PreferSimpleProtocol: pgprestatement }), &gorm.Config{})
	if err != nil {
		return nil,err
	}
	sqlDB,err := db.DB()
	if err != nil {
		return nil,err
	}
	sqlDB.SetMaxIdleConns(cfg.PGMaxIdelConns)
	sqlDB.SetMaxOpenConns(cfg.PGMaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.PGConnMaxLifetime)*time.Minute)
	return db,nil
}

func  DBClose(){
	sqlDB,err := PG.DB()
	if err != nil {
		zap.L().Error("func DBClose: ",zap.Error(err))
	}
	sqlDB.Close()
}