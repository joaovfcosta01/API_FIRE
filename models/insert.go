package models

import (
	"API_FIRE/configs"
	"database/sql"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := sql.Open("firebirdsql", configs.GetDBDSN())
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES (?, ?, ?) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)
	return
}
