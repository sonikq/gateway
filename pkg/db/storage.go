package db

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	cfg "gitlab.geogracom.com/skdf/skdf-abac-go/configs/app"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db/admin"
	user "gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db/user"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db/user_access"
	"time"
)

const (
	driverName = "postgres"
)

type DB struct {
	user.IUserDB
	user_access.IAccessUserDB
	admin.IAdminDB
}

type connectResponse struct {
	db  *sqlx.DB
	err error
}

func ConnectContext(ctx context.Context, config cfg.PostgresConfig) (*DB, error) {
	ctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()

	ch := make(chan connectResponse)

	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)

	go func() {
		db, err := connect(driverName, connInfo)
		ch <- connectResponse{
			db:  db,
			err: err,
		}
	}()

	for {
		select {
		case resp := <-ch:
			return &DB{
				IUserDB:       user.NewUserDB(resp.db),
				IAccessUserDB: user_access.NewAccessUserDB(resp.db),
				IAdminDB:      admin.NewAdminDB(resp.db),
			}, resp.err
		case <-ctx.Done():
			return nil, fmt.Errorf("connecting to database took long time")
		}
	}

}

func connect(driverName string, connInfo string) (_db *sqlx.DB, err error) {
	_db, err = sqlx.Connect(driverName, connInfo)
	return
}

func (db *DB) Close() error {
	if err := db.IUserDB.CloseDB(); err != nil {
		return err
	}

	if err := db.IAccessUserDB.CloseDB(); err != nil {
		return err
	}

	return nil
}
