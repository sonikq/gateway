package admin

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"time"
)

type AdminDB struct {
	*sqlx.DB
}

type IAdminDB interface {
	AddUserMatrix(newUserID, newObjectValue int, newObjectType, newPeriodStart, newRole, newPeriodFinish string, newStates []string, newOrgsID []int) error
	UpdateUserMatrix(newUserID, newObjectValue int, newObjectType, newPeriodStart, newRole, newPeriodFinish string, newStates []string, newOrgsID []int) error
	DeleteUserMatrix(oldUserID int) error
}

func NewAdminDB(db *sqlx.DB) *AdminDB {
	return &AdminDB{
		DB: db,
	}
}

func (db *AdminDB) AddUserMatrix(newUserID, newObjectValue int, newObjectType, newPeriodStart, newRole, newPeriodFinish string, newStates []string, newOrgsID []int) error {
	start := time.Now()
	//periodStart, _ := time.Parse("2006-01-02", newPeriodStart)
	//periodFinish, _ := time.Parse("2006-01-02", newPeriodFinish)
	_, err := db.Exec(addUserMatrixByAdmin, newUserID, newObjectValue, newObjectType, newPeriodStart, newRole, newPeriodFinish, pq.Array(newStates), pq.Array(newOrgsID))
	if err != nil {
		return err
	}
	log.Printf("Adding user by admin time: %v\n", time.Since(start))
	return nil
}

func (db *AdminDB) UpdateUserMatrix(newUserID, newObjectValue int, newObjectType, newPeriodStart, newRole, newPeriodFinish string, newStates []string, newOrgsID []int) error {
	start := time.Now()
	periodStart, _ := time.Parse("2006-01-02", newPeriodStart)
	periodFinish, _ := time.Parse("2006-01-02", newPeriodFinish)
	_, err := db.Exec(updateUserMatrixByAdmin, newUserID, newObjectValue, newObjectType, periodStart, newRole, periodFinish, pq.Array(newStates), pq.Array(newOrgsID))
	if err != nil {
		return err
	}
	log.Printf("Updating user by admin time: %v\n", time.Since(start))
	return nil
}

func (db *AdminDB) DeleteUserMatrix(oldUserID int) error {
	_, err := db.Exec(deleteUserMatrixByAdmin, oldUserID)
	if err != nil {
		return err
	}

	return nil
}
