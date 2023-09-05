package user_access

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"time"
)

type AccessUserDB struct {
	*sqlx.DB
}

type IAccessUserDB interface {
	AddUserAccess(int, []int, []int, []int, []string) error
	CheckUserAccess(int, int, int, int, string) (bool, error)
	DeleteUserAccess(int) error
	UpdateUserAccess(int, []int, []int, []int, []string) error
	CloseDB() error
}

func NewAccessUserDB(db *sqlx.DB) *AccessUserDB {
	return &AccessUserDB{
		DB: db,
	}
}

func (db *AccessUserDB) AddUserAccess(newUserId int, newOkuds []int, newReportIDs []int, newChapterIDs []int, newOperations []string) error {
	start := time.Now()
	_, err := db.Exec(addUserAccess, newUserId, pq.Array(newOkuds), pq.Array(newReportIDs), pq.Array(newChapterIDs), pq.Array(newOperations))
	if err != nil {
		return err
	}
	log.Printf("Adding user access time: %v\n", time.Since(start))
	return nil
}

func (db *AccessUserDB) CheckUserAccess(userID int, okud int, reportID int, chapterID int, operation string) (bool, error) {
	var exists bool
	row, err := db.Query(checkUserAccess, userID, okud, reportID, chapterID, operation)
	if err != nil {
		return false, err
	}
	defer row.Close()

	for row.Next() {
		if err = row.Scan(&exists); err != nil {
			return false, err
		}
	}

	return exists, nil
}

func (db *AccessUserDB) DeleteUserAccess(userID int) error {
	_, err := db.Exec(deleteUserAccess, userID)
	if err != nil {
		return err
	}

	return nil
}

func (db *AccessUserDB) UpdateUserAccess(newUserId int, newOkuds []int, newReportIDs []int, newChapterIDs []int, newOperations []string) error {
	start := time.Now()
	_, err := db.Exec(updateUserAccess, newUserId, pq.Array(newOkuds), pq.Array(newReportIDs), pq.Array(newChapterIDs), pq.Array(newOperations))
	if err != nil {
		return err
	}
	log.Printf("Updating user access time: %v\n", time.Since(start))
	return nil
}

func (db *AccessUserDB) CloseDB() error {
	return db.Close()
}
