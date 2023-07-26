package mysql

import (
	"errors"
	"go.uber.org/zap"
	"multiovirt-admin/models"
	"multiovirt-admin/pkg/encryptpassword"
)

func CheckOvirtDBConf(o *models.OvirtConf) (err error) {
	var count int64
	result := Mysql.Raw(`select count(*) from  pg_info where host = ? and port = ? and dbname = ?`, o.Host, o.Port, o.DBName).Count(&count)
	//err = Mysql.Raw(`select count(*) from  pg_info where host = ? and port = ? and dbname = ?`,o.Host,o.Port,o.DBName).Count(&count).Error
	if result.Error != nil {
		zap.L().Error("CheckOvirtDBConf error", zap.Error(err))
		return err
	}
	if count > 0 {
		return errors.New("CheckOvirtDBConf config exist!")
	} else {
		zap.L().Info("CheckOvirtDBConf pass!")
		return nil
	}
}

func InsertOvirtDBConf(o *models.OvirtConf) (rowAffected int64, err error) {
	o.Password = encryptpassword.EncryptPassword(o.Password)
	result := Mysql.Exec(`insert into pg_info(host,port,username,password,dbname,timezone,pgmaxidelconns,pgmaxopenconns,pgconnmaxlifetime,pgprestatement)
	           value(?,?,?,?,?,?,?,?,?,?)`, o.Host, o.Port, o.User, o.Password, o.DBName, o.TimeZone, o.PGMaxIdelConns, o.PGMaxOpenConns, o.PGConnMaxLifetime, o.PGPreStatement)
	rowAffected = result.RowsAffected
	if rowAffected == 0 {
		zap.L().Error("InsertOvirtDBConf no data insert!")
		return 0, errors.New("InsertOvirtDBConf no data insert!")
	} else {
		zap.L().Info("InsertOvirtDBConf insert success!")
		return rowAffected, nil
	}
}

func ListOvirt() (listovirt []*models.ListOvirtConf, err error) {
	result := Mysql.Raw("select aliasname,host,port from pg_info").Scan(&listovirt)
	if result.Error != nil {
		zap.L().Error("ListOvirt Error!")
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		zap.L().Warn("ListOvirt Record Not Found!")
		return nil, errors.New("ListOvirt Record Not Found!")
	}
	return listovirt, nil
}

func GetOvirtConfDetail(aliasname string) (listovirt []*models.GetOvirtConfDetail, err error) {
	result := Mysql.Raw("select id,aliasname,host,port,dbname,timezone,pgmaxidelconns,pgmaxopenconns,pgconnmaxlifetime,pgprestatement from pg_info where aliasname = ?", aliasname).Scan(&listovirt)
	if result.Error != nil {
		zap.L().Error("ListOvirt Error!")
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		zap.L().Warn("ListOvirt Record Not Found!")
		return nil, errors.New("ListOvirt Record Not Found!")
	}
	return listovirt, nil
}
