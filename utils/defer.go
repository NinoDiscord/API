package utils

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"io"
)

func SwallowError(rows *sql.Rows) {
	err := rows.Close(); if err != nil {
		logrus.Fatal(err)
	}
}

func SwallowHttpError(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		logrus.Fatalf("Unable to close response body: %v", err)
	}
}
