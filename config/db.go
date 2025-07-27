package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db   *sql.DB
	once sync.Once
)

func SetupDatabase() *sql.DB {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Erro ao carregar .env")
		}

		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USERNAME")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbUser, dbPassword, dbHost, dbPort, dbName,
		)

		db, err = sql.Open("mysql", connectionStr)
		if err != nil {
			log.Fatal("Erro ao abrir conexão com o banco:", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal("Erro ao conectar ao banco:", err)
		}

		fmt.Println("✅ Conectado ao banco MySQL com sucesso!")
	})

	return db
}
