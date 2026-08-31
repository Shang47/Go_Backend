package poker

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewMariaPlayerStore() *MariaPlayerStore {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	cfg.DBName = os.Getenv("DB_NAME")

	// Get a database handle.
	var db *sql.DB
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	//fmt.Println("Connected")

	return &MariaPlayerStore{db}
}

type score struct {
	name           string
	personal_score int
}
type MariaPlayerStore struct {
	DBHandle *sql.DB
}

func (m *MariaPlayerStore) GetPlayerScore(name string) int {
	var scr score

	row := m.DBHandle.QueryRow("SELECT * FROM scores WHERE name = ?", name)
	if err := row.Scan(&scr.name, &scr.personal_score); err != nil {
		if err == sql.ErrNoRows {
			log.Println(fmt.Errorf("Query error: No data with name %s", name))
			return 0
		}
		log.Println(fmt.Errorf("Query error: Encounter error (%v) when querying for name %s", err, name))
		return 0
	}
	return scr.personal_score
}
func (m *MariaPlayerStore) RecordWin(name string) {
	var prevScore int
	prevScore = m.GetPlayerScore(name)

	if prevScore == 0 {
		_, err := m.DBHandle.Exec("INSERT INTO scores (name, win) VALUES (?, ?)", name, 1)
		if err != nil {
			fmt.Printf("Insert error: Encounter error (%v) when inserting data with name %s\n", err, name)
		}
		fmt.Printf("Insert completed: Successfully insert data with name %s\n", name)
	} else {
		_, err := m.DBHandle.Exec("UPDATE scores SET win = ? WHERE name = ?", prevScore+1, name)
		if err != nil {
			fmt.Printf("Insert error: Encounter error (%v) when inserting data with name %s\n", err, name)
		}
	}
}
func (m *MariaPlayerStore) GetLeague() League {
	return nil
}
func (m *MariaPlayerStore) close() {
	err := m.DBHandle.Close()
	if err != nil {
		log.Printf("Encounter error when trying to close DB: %v\n", err)
	}
}
