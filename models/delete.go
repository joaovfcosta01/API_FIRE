package models

import (
	"API_FIRE/configs"
	"database/sql"
)

func Delete(id int64) (int64, error) {
	conn, err := sql.Open("firebirdsql", configs.GetDBDSN())
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=?`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
