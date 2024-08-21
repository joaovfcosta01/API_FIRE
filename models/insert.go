package models

import "API_FIRE/db"

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection1()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES (?, ?, ?) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)
	return
}
