package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type User struct {
	ID           int64
	Name         string
	Email        string
	Password     string
	RegisteredAt time.Time
}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres"+
		" sslmode=disable password=7777")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	err = insertUser(User{
		Name:     "Leonidas",
		Email:    "Random@mail.ru",
		Password: "Easypassword",
	}, db)

	fmt.Println(mid(getUsers(db)))
}

func mid(u []User, e error) []User {
	_ = e
	return u
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := new(User)
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func getUserById(id int64, db *sql.DB) (User, error) {
	u := new(User)
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)

	return *u, err
}

/*
This function inserts new user in database
*/
func insertUser(u User, db *sql.DB) error {
	tx, err := db.Begin() // transaction -> or all actions should be completed or none of them
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO logs (entity, action) VALUES ($1, $2)", "user", "created")
	if err != nil {
		return err
	}
	return tx.Commit()
}

func deleteUser(id int64, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func updateUser(u User, db *sql.DB, id int64) error {
	_, err := db.Exec("update users set name = $1, email = $2 where id = $3", u.Name, u.Email, id)
	return err
}
