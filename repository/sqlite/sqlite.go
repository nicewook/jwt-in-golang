package sqlite

import (
	"database/sql"
	"fmt"
	"jwt-in-golang/entity"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const userDB = "./user.db"

// https://manufacture.tistory.com/17
type SQLiteDB struct {
	db *sql.DB
}

func createTable(db *sql.DB) error {
	usersTableSQL := `CREATE TABLE IF NOT EXISTS users (
		username   TEXT,
		password   TEXT,
		first_name TEXT,
		last_name  TEXT,
		address    TEXT,
		email      TEXT
	);`

	log.Println("Create transaction users table...")

	_, err := db.Exec(usersTableSQL) // Execute SQL Statements
	if err != nil {
		log.Printf("fail to db.Exec(): %v", err)
		return err
	}
	return nil
}

func NewSQLiteDB() *SQLiteDB {
	db, err := sql.Open("sqlite3", userDB) // Open the created SQLite File
	if err != nil {
		log.Fatal("fail to open", userDB)
	}

	// create table and index
	if err := createTable(db); err != nil { // Create Database Tables
		log.Println("fail to createTable: ", err)
		log.Fatal("fail to open", userDB)
	}

	return &SQLiteDB{
		db: db,
	}
}

func (s *SQLiteDB) GetUser(username string) (entity.User, error) {
	var user entity.User

	stmt, err := s.db.Prepare("SELECT * FROM users WHERE username=?")
	if err != nil {
		log.Println(err)
		// return user, err
	}
	defer stmt.Close()

	if err := stmt.QueryRow(username).
		Scan(&user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Address, &user.Email); err != nil {
		log.Println(err)
		// return user, err
	}
	fmt.Println(user)
	return user, nil
}

func (s *SQLiteDB) AddUser(user entity.User) error {
	_, err := s.db.Exec("INSERT INTO users (username, password, first_name, last_name, address, email) VALUES (?,?,?,?,?,?)",
		user.Username, user.Password, user.FirstName, user.LastName, user.Address, user.Email)
	return err
}
