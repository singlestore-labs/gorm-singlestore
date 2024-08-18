package singlestore

import (
	"github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
)

// The error codes to map singlestore and mysql errors to gorm errors.
// Here is the singlestore error codes reference https://docs.singlestore.com/cloud/reference/troubleshooting-reference/query-errors/
// Here is the mysql error codes reference https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html.
var errCodes = map[uint16]error{
	1062: gorm.ErrDuplicatedKey,
}

func (dialector Dialector) Translate(err error) error {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if translatedErr, found := errCodes[mysqlErr.Number]; found {
			return translatedErr
		}
		return mysqlErr
	}

	return err
}
