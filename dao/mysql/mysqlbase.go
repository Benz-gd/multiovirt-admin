package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"multiovirt-admin/settings"
	"time"
)

var MysqlBase *gorm.DB

func Init(cfg *settings.MySQLBase) (*gorm.DB, error) {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
	//	viper.GetString("mysql.user"),
	//	viper.GetString("mysql.password"),
	//	viper.GetString("mysql.host"),
	//	viper.GetInt("mysql.port"),
	//	viper.GetString("mysql.dbname"),
	//	viper.GetString("mysql.mysqlcharset"),
	//	viper.GetString("mysql.mysqlcollation"),
	//	viper.GetString("mysql.mysqlquery"),
	//)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.MysqlCharset,
		cfg.MysqlCollation,
		cfg.MysqlQuery,
	)
	//fmt.Printf("dsn is:%s\n",dsn)
	var err error
	MysqlBase, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		zap.L().Error("func initMysqlClient: ", zap.Error(err))
		return nil, err
	}

	sqlDB, _ := MysqlBase.DB()
	sqlDB.SetMaxIdleConns(cfg.MysqlMaxIdelConns)
	sqlDB.SetMaxOpenConns(cfg.MysqlMaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MysqlConnMaxLifetime) * time.Minute)
	return MysqlBase, nil
}

func DBClose() {
	sqlDB, err := MysqlBase.DB()
	if err != nil {
		zap.L().Error("func DBClose: ", zap.Error(err))
	}
	sqlDB.Close()
}
