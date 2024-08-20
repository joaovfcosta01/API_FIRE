package models

import (
	"API_FIRE/configs"
	"database/sql"
)

func Update(id int64, todo Todo) (int64, error) {
	conn, err := sql.Open("firebirdsql", configs.GetDBDSN())
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(
		`UPDATE todos SET title=?, description=?, done=? WHERE id=?`,
		todo.Title, todo.Description, todo.Done, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
