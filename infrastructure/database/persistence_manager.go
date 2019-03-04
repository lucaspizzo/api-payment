package database

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const defaultLifeTime = time.Minute * 5
const defaultMaxIdleConns = 5
const defaultMaxOpenConns = 5

type MultipleDatabaseOnPoolError struct{}

func (e *MultipleDatabaseOnPoolError) Error() string {
	return fmt.Sprintf("It isn't allowed define a default database. You should pass the database name instead.")
}

type PoolWithoutInstanceError struct{}

func (e *PoolWithoutInstanceError) Error() string {
	return fmt.Sprintf("Can't define a default database. You should Set or Load a instance first.")
}

func defineDatabaseName(dbNameParam []string, poolSize int, firstKeyOfPool func() string) (string, error) {
	var dbName string

	if len(dbNameParam) == 0 {
		if poolSize == 0 {
			return "", &PoolWithoutInstanceError{}
		}

		if poolSize > 1 {
			return "", &MultipleDatabaseOnPoolError{}
		}

		dbName = firstKeyOfPool()
	} else {
		dbName = dbNameParam[0]
	}

	return dbName, nil
}

func generateDsn(driverName string, user string, pass string, host string, port int, dbName string, sslMode bool) (string, error) {
	if driverName == "mysql" {
		return user + ":" + pass + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbName + "?charset=utf8&parseTime=True&loc=America%2FSao_Paulo", nil
	}

	if driverName == "postgres" {

		var sslOption string

		if sslMode {
			sslOption = "enable"
		} else {
			sslOption = "disable"
		}

		return "host=" + host +
			" port=" + strconv.Itoa(port) +
			" user=" + user +
			" dbname=" + dbName +
			" password=" + pass +
			" sslmode=" + sslOption, nil
	}

	return "", errors.New("Can't generate DSN for " + driverName)
}
