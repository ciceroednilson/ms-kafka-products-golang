package productrepository

import (
	"database/sql"
	"fmt"
	"go/core/ports"
	"go/infrastructure/entities"
	"go/util/mysql"
	"log"
)

const selectAllProducts = "SELECT id, description, price, total, created FROM tb_product"
const selectByIdProduct = "SELECT id, description, price, total, created FROM tb_product WHERE id = ?"
const deleteByIdProduct = "DELETE FROM tb_product WHERE id=?"
const insertProduct = "INSERT INTO tb_product (description, price, total, created) VALUES (?, ?, ?, ?)"
const updateProduct = "UPDATE tb_product SET description=?, price=?,total=? WHERE id=?"

type Repository struct{}

func NewRepository() ports.ProductRepositoryPort {
	return &Repository{}
}

func (r *Repository) FindAll() (*[]entities.Product, error) {
	db := mysql.Open()
	defer db.Close()
	var products []entities.Product
	rows, err := db.Query(selectAllProducts)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	for rows.Next() {
		product := entities.Product{}
		if err := rows.Scan(&product.Id, &product.Description, &product.Price, &product.Total, &product.Created); err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	rows.Close()
	return &products, nil
}

func (r *Repository) FindById(key int) (*entities.Product, error) {
	var product entities.Product
	db := mysql.Open()
	defer db.Close()
	row := db.QueryRow(selectByIdProduct, key)
	err := row.Scan(&product.Id, &product.Description, &product.Price, &product.Total, &product.Created)
	if err == nil {
		return &product, nil
	}
	switch err {
	case sql.ErrNoRows:
		fmt.Printf("register not found to product id %d", key)
		return nil, nil
	default:
		return nil, err
	}
}

func (r *Repository) DeleteById(key int) error {
	db := mysql.Open()
	defer db.Close()
	_, err := db.Exec(deleteByIdProduct, key)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *Repository) Save(product entities.Product) error {
	db := mysql.Open()
	defer db.Close()
	_, err := db.Exec(insertProduct, product.Description, product.Price, product.Total, product.Created)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *Repository) Update(product entities.Product) error {
	db := mysql.Open()
	defer db.Close()
	_, err := db.Exec(updateProduct, product.Description, product.Price, product.Total, product.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
