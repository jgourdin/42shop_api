package model

import (
	"database/sql"
	"fmt"
)

type product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
	Type        string  `json:"type"`
}

type user struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Mail      string `json:"mail"`
	Commandes string `json:"commandes"`
}

//Query Row return one Row, Query only one

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price, description, stock, image, type FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price, &p.Description, &p.Stock, &p.Image, &p.Type)
}
func (p *product) updateProduct(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE products SET name=$1, price=$2, description=$3, stock=$4, image=$5, type=$6 WHERE id=$7",
			p.Name, p.Price, p.Description, p.Stock, p.Image, p.Type, p.ID)
	return err
}
func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)
	return err
}
func (p *product) createProduct(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO products(name, price, description, stock, image, type) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		p.Name, p.Price, p.Description, p.Stock, p.Image, p.Type).Scan(&p.ID)
	if err != nil {
		return err
	}
	return nil
}

//defer is a callback

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name, price, description, stock, image, type FROM products LIMIT $1 OFFSET $2",
		count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []product{}
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Stock, &p.Image, &p.Type); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	fmt.Println(products)
	return products, nil
}

/* USER
Name      string `json:"name"`
Password  string `json:"password"`
Mail      string `json:"mail"`
Commandes string `json:"commandes"`
*/

func (p *user) getUser(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM user WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Password, &p.Mail, &p.Commandes)
}
func (p *user) updateUser(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE user SET name=$1, password=$2, mail=$3, commandes=$4, WHERE id=$4",
			p.Name, p.Password, p.Mail, p.Commandes, p.ID)
	return err
}
func (p *user) deleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM user WHERE id=$1", p.ID)
	return err
}
func (p *user) createUser(db *sql.DB) error {
	fmt.Println("CREATE USER")
	fmt.Println(p)
	err := db.QueryRow(
		"INSERT INTO user(name, password, mail, commandes) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		p.Name, p.Password, p.Mail, p.Commandes).Scan(&p.ID)
	if err != nil {
		return err
	}
	return nil
}

//defer is a callback

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	rows, err := db.Query(
		"SELECT id, name, password, mail, commandes, type FROM user LIMIT $1 OFFSET $2",
		count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []user{}
	for rows.Next() {
		var p user
		if err := rows.Scan(&p.ID, &p.Name, &p.Password, &p.Mail, &p.Commandes); err != nil {
			return nil, err
		}
		users = append(users, p)
	}
	return users, nil
}
