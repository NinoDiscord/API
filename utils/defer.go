package utils

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func SwallowError(rows *sql.Rows) {
	err := rows.Close(); if err != nil {
		logrus.Fatal(err)
	}
}
