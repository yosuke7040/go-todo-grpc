//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repositories

import (
	"database/sql"
	"log"

	"github.com/yosuke7040/go-todo-grpc/models"
)

type TodoRepoInterface interface {
	SelectTodoList() ([]*models.Todo, error)
	SelectTodo(int32) (*models.Todo, error)
	InsertTodo(string) error
	UpdateTodo(int32, string, int32) error
	DeleteTodo(int32) error
	IsValidTodoId(int32) (bool, error)
}

type Todo struct {
	db *sql.DB
}

func NewTodo(db *sql.DB) *Todo {
	return &Todo{db: db}
}

// taskの一覧を取得する
func (t *Todo) SelectTodoList() ([]*models.Todo, error) {
	const sqlStr = `SELECT id, title, status FROM todos`

	rows, err := t.db.Query(sqlStr)
	if err != nil {
		log.Printf("SelectTodoList error: %v", err)
		return nil, err
	}
	defer rows.Close()

	todos := make([]*models.Todo, 0)
	for rows.Next() {
		todo := &models.Todo{}
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Status); err != nil {
			log.Printf("SelectTodoList error: %v", err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// 指定したIDのtaskを取得する
func (t *Todo) SelectTodo(id int32) (*models.Todo, error) {
	const sqlStr = `SELECT id, title, status FROM todos WHERE id = ?`
	// NOTE: ↓の場合は、todoがnilのままになるのでScanできない
	// var todo *models.Todo
	todo := &models.Todo{}

	err := t.db.QueryRow(sqlStr, id).Scan(&todo.Id, &todo.Title, &todo.Status)
	if err != nil {
		log.Printf("SelectTodoList error: %v", err)
		return nil, err
	}

	return todo, nil
}

// taskを作成する
func (t *Todo) InsertTodo(title string) error {
	const sqlStr = `INSERT INTO todos (title, status) VALUES (?, ?)`

	_, err := t.db.Exec(sqlStr, title, models.TodoStatusUnspecified)
	if err != nil {
		log.Printf("InsertTodo error: %v", err)
		return err
	}

	return nil
}

// taskを更新する
func (t *Todo) UpdateTodo(id int32, title string, status int32) error {
	const sqlStr = `UPDATE todos SET title = ?, status = ? WHERE id = ?`

	_, err := t.db.Exec(sqlStr, title, status, id)
	if err != nil {
		log.Printf("UpdateTodo error: %v", err)
		return err
	}

	return nil
}

// taskを削除する
func (t *Todo) DeleteTodo(id int32) error {
	const sqlStr = `DELETE FROM todos WHERE id = ?`

	_, err := t.db.Exec(sqlStr, id)
	if err != nil {
		log.Printf("DeleteTodo error: %v", err)
		return err
	}

	return nil
}

// 存在するIDかどうかを確認する
func (t *Todo) IsValidTodoId(id int32) (bool, error) {
	isValid := false
	const sqlstr = `SELECT exists (
		SELECT 1 FROM todos WHERE id = ?
	)`

	err := t.db.QueryRow(sqlstr, id).Scan(&isValid)
	if err != nil {
		log.Printf("IsValidTodoId error: %v", err)
	}
	return isValid, err
}
