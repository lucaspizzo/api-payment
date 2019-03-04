package database

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type singleton struct {
	Connection *gorm.DB
}

var (
	poolGormDb    = make(map[string]*gorm.DB)
	LogModeEnable = func() bool {
		return true
	}()
)

func LoadGormPostGres(user string, pass string, host string, port int, dbName string, sslMode bool) error {
	return LoadGorm("postgres", user, pass, host, port, dbName, sslMode)
}

func LoadGorm(driverName string, user string, pass string, host string, port int, dbName string, sslMode bool) error {
	var err error

	if poolGormDb[dbName] == nil {
		poolGormDb[dbName], err = getGormConnection(driverName, user, pass, host, port, dbName, sslMode)
	}

	return err
}


func GetGormDb(dbNameParam ...string) (*gorm.DB, error) {
	dbName, err := defineDatabaseName(dbNameParam, len(poolGormDb), func() string {
		return firstKeyFromGormPool(poolGormDb)
	})

	if err != nil {
		return nil, err
	}
	if poolGormDb[dbName] == nil {
		return nil, errors.New("LoadGorm/SetGormDb wasn't called for database: " + dbName)
	}

	return poolGormDb[dbName], nil
}

func SetGormDb(gormDb *gorm.DB, dbName string) {
	poolGormDb[dbName] = gormDb
}

func getGormConnection(driverName string, user string, pass string, host string, port int, dbName string, sslMode bool) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var dsn string
	dsn, err = generateDsn(driverName, user, pass, host, port, dbName, sslMode)
	db, err = gorm.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err)
	}
	db.DB().SetConnMaxLifetime(defaultLifeTime)
	db.DB().SetMaxIdleConns(defaultMaxIdleConns)
	db.DB().SetMaxOpenConns(defaultMaxOpenConns)
	db.LogMode(LogModeEnable)
	return db, err
}

func firstKeyFromGormPool(object map[string]*gorm.DB) string {
	for k := range object {
		return k
	}

	return ""
}

func PurgeGormPool() {
	for k, v := range poolGormDb {
		v.Close()
		delete(poolGormDb, k)
	}
}