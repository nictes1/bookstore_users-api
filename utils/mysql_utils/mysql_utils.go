package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/nictes1/bookstore_users-api/utils/errors"

	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("not record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalide data")
	}
	return errors.NewInternalServerError("error processing request")
}
