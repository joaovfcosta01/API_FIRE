package models

import (
	"API_FIRE/configs"
	"database/sql"
)

func GetAll() (todos []Todo, err error) {
	conn, err := sql.Open("firebirdsql", configs.GetDBDSN())
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}
	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}
	return
}
