package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

type UserInfo struct {
	UserId       int
	Username     string
	CountCorrect int
	CountFalse   int
	Points       int
}

func NewDatabase(path string) (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{db: db}, nil
}

func (d *Database) Init(tableName string) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (user_id INT, username TEXT, count_correct INT, count_false INT, points INT)", tableName)
	if _, err := d.db.ExecContext(context.Background(), query); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}

func (d *Database) AddUser(tableName string, userId int, username string) error {
	query := fmt.Sprintf("INSERT INTO %v (user_id, username, count_correct, count_false, points) VALUES (?, ?, 0, 0, 0)", tableName)
	if _, err := d.db.ExecContext(context.Background(), query, userId, username); err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}
	return nil
}

func (d *Database) CheckUserExists(tableName string, userId int, username string) error {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %v WHERE user_id = ?)", tableName)
	var exists bool
	err := d.db.QueryRow(query, userId).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if exists {
		query = fmt.Sprintf("SELECT user_id, username, points, count_correct, count_false FROM %v", tableName)
		var user UserInfo
		err = d.db.QueryRow(query, userId).Scan(&user.UserId, &user.Username, &user.CountCorrect, &user.CountFalse, &user.Points)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}
		if user.Username == username {
			return nil
		}
		// todo add change username
		return nil
	}

	err = d.AddUser(tableName, userId, username)
	if err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}

	return nil
}

func (d *Database) IncreaseValue(tableName string, columnName string, userId int, value int) error {
	query := fmt.Sprintf("UPDATE %v SET %v = %v + ? WHERE user_id = ?", tableName, columnName, columnName)
	if _, err := d.db.ExecContext(context.Background(), query, value, userId); err != nil {
		return fmt.Errorf("failed to increase value: %w", err)
	}
	return nil
}

func (d *Database) GetTopUsers(tableName string, usersCount int) ([]UserInfo, error) {
	var query string
	if usersCount == -1 {
		query = fmt.Sprintf("SELECT * FROM %v", tableName)
	} else {
		query = fmt.Sprintf("SELECT * FROM %v ORDER BY count_correct DESC LIMIT %v", tableName, usersCount)
	}
	rows, err := d.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get top users: %w", err)
	}
	var users []UserInfo
	for rows.Next() {
		var user UserInfo
		err = rows.Scan(&user.UserId, &user.Username, &user.CountCorrect, &user.CountFalse, &user.Points)
		if err != nil {
			return nil, fmt.Errorf("failed to get top users: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}
