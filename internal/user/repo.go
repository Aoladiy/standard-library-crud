package user

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

type Repo struct {
	db *sql.DB
}

func (r Repo) beginTransaction() (tx *sql.Tx, rollback func(tx *sql.Tx)) {
	tx, err := r.db.Begin()
	rollback = func(tx *sql.Tx) {
		_ = tx.Rollback()
	}
	if err != nil {
		log.Println("some trouble during transaction beginning", err)
		return nil, rollback
	}
	return tx, rollback
}

func (r Repo) getUserById(id int) (user User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()
	row := r.db.QueryRowContext(ctx, "SELECT id, email, fullname, phonenumber, age FROM users WHERE id = $1", id)
	err = row.Scan(&user.Id, &user.Email, &user.FullName, &user.PhoneNumber, &user.Age)
	if err != nil {
		log.Println("some trouble while scanning users by id", err)
		return User{}, err
	}
	return user, nil
}

func (r Repo) getUsers() ([]User, error) {
	users := make([]User, 0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()
	rows, err := r.db.QueryContext(ctx, "SELECT id, email, fullname, phonenumber, age FROM users")
	if err != nil {
		log.Println("some trouble while selecting all users", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Email, &user.FullName, &user.PhoneNumber, &user.Age)
		if err != nil {
			log.Println("some trouble while scanning all users", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r Repo) createUser(user User) (id int, err error) {
	tx, rollback := r.beginTransaction()
	defer rollback(tx)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()
	row := tx.QueryRowContext(
		ctx,
		`INSERT INTO users (email, fullname, phonenumber, age) VALUES ($1, $2, $3, $4) returning id`,
		user.Email,
		user.FullName,
		user.PhoneNumber,
		user.Age)
	err = row.Scan(&id)
	if err != nil {
		log.Println("some trouble while scanning id", err)
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("some trouble while commiting transaction", err)
		return 0, err
	}
	return id, nil
}

func (r Repo) updateUser(user User) (err error, ok bool) {
	tx, rollback := r.beginTransaction()
	defer rollback(tx)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()
	result, err := r.db.ExecContext(ctx,
		"UPDATE users SET email = $1, fullname = $2, phonenumber = $3, age = $4 where id = $5",
		user.Email, user.FullName, user.PhoneNumber, user.Age, user.Id)
	if err != nil {
		log.Println("some trouble while updating user by id", err)
		return err, false
	}
	err = tx.Commit()
	if err != nil {
		log.Println("some trouble while commiting transaction", err)
		return err, false
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Println("some trouble while deleting user by id", err)
		return err, false
	}
	if affected < 1 {
		return errors.New("everything ok, but nothing deleted. Most likely there is just no row with such id"), true
	}
	return nil, true
}

func (r Repo) deleteUserById(id int) (err error, ok bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()
	result, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Println("some trouble while deleting user by id", err)
		return err, false
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Println("some trouble while deleting user by id", err)
		return err, false
	}
	if affected < 1 {
		return errors.New("everything ok, but nothing deleted. Most likely there is just no row with such id"), true
	}
	return nil, false
}
