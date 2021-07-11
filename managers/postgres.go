package managers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"nino.sh/api/utils"
	"os"
	"strconv"
	"time"
)

type PostgresManager struct {
	Connection *sql.DB
}

func NewPostgresManager() *PostgresManager {
	return &PostgresManager{
		Connection: nil,
	}
}

func (pg *PostgresManager) GetConnection() error {
	port, err := strconv.Atoi(os.Getenv("PG_PORT")); if err != nil {
		return err
	}

	host := os.Getenv("PG_HOST"); if host == "" {
		return errors.New("host must be defined in environment variables")
	}

	username := os.Getenv("PG_USERNAME"); if username == "" {
		return errors.New("username must be defined in environment variables")
	}

	password := os.Getenv("PG_PASSWORD"); if password == "" {
		return errors.New("password must be defined in environment variables")
	}

	db := os.Getenv("PG_DATABASE"); if db == "" {
		db = "nino"
	}

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, db)
	conn, err := sql.Open("postgres", sqlInfo); if err != nil {
		return err
	}

	err = conn.PingContext(context.TODO()); if err != nil {
		return err
	}

	logrus.Info("Connected to PostgreSQL. :3")
	pg.Connection = conn

	return nil
}

// GetPing returns the latency from calculating from selecting all guilds.
func (db *PostgresManager) GetPing() int64 {
	ping := time.Now()
	rows, err := db.Connection.QueryContext(context.TODO(), "SELECT * FROM guilds;"); if err != nil {
		return -1
	}

	defer utils.SwallowError(rows)
	return time.Since(ping).Milliseconds()
}
