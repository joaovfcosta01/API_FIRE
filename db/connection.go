package db

import (
	"API_FIRE/configs"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/nakagami/firebirdsql" // Importa o driver Firebird
)

// OpenConnection abre uma conexão com o banco de dados Firebird
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	// Formatar a string de conexão de acordo com o driver Firebird
	sc := fmt.Sprintf(
		"localhost/%s:%s?user=%s&password=%s",
		conf.Port,
		conf.Database,
		conf.User,
		conf.Pass,
	)

	conn, err := sql.Open("firebirdsql", sc) // Usar o nome do driver correto
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
func OpenConnection1() (*sql.DB, error) {
	conf := configs.GetDB()
	// Montando a string de conexão
	sc := fmt.Sprintf("%s:%s@%s:%s/%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)
	// Abrindo a conexão
	println(sc)
	conn, err := sql.Open("firebirdsql", sc)

	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()

	// Testando a conexão
	err = conn.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	return conn, err

}
