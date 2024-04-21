package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccount(int) (*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr:= "user=postgres dbname=postgres password=mypassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db}, nil
}

func (s *PostgresStorage) Init() error {
	return s.CreateTable()
	
}

func (s *PostgresStorage) CreateTable() error {
	query:= `CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		number_bank BIGINT NOT NULL,
		balance FLOAT NOT NULL,
		created_at TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateAccount(a *Account) error {
	query:= `INSERT INTO accounts (first_name, last_name, number_bank, balance, created_at) VALUES ($1, $2, $3, $4, $5)`
	resp, err:= s.db.Query(query, a.FirstName, a.LastName, a.NumberBank, a.Balance, a.CreateAt)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil	
}

func (s *PostgresStorage) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStorage) GetAccount(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStorage) GetAccounts() ([]*Account, error) {
	rows, err:= s.db.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts:= []*Account{}

	for rows.Next() {
		account:= &Account{}
		if err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.NumberBank, &account.Balance, &account.CreateAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil

}