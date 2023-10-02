//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repositories

import (
	"database/sql"
	"log"

	"github.com/yosuke7040/go-todo-grpc/models"
)

type UserRepositoryInterface interface {
	SelectUser(int32) (*models.User, error)
	IsValidUserId(int32) (bool, error)
}

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

// 指定したIDのuserを取得する
func (u *User) SelectUser(id int32) (*models.User, error) {
	const sqlStr = `SELECT id, name, email FROM users WHERE id = ?`
	// NOTE: ↓の場合は、userがnilのままになるのでScanできない
	// var user *models.User
	user := &models.User{}

	err := u.db.QueryRow(sqlStr, id).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// 存在するIDかどうかを確認する
func (t *User) IsValidUserId(id int32) (bool, error) {
	isValid := false
	const sqlstr = `SELECT exists (
		SELECT 1 FROM users WHERE id = ?
	)`

	err := t.db.QueryRow(sqlstr, id).Scan(&isValid)
	if err != nil {
		log.Printf("IsValidUserId error: %v", err)
	}
	return isValid, err
}
