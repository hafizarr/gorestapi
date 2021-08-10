package warehouse

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"gorestapi/config"
	"gorestapi/models"
)

const (
	table = "products"
)

// GetAll
func GetAll(ctx context.Context, sortBy string) ([]models.Product, error) {

	var products []models.Product

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := ""
	param := sortBy
	switch param {
	case "lowest-price":
		queryText = fmt.Sprintf("SELECT * FROM %v Order By product_price ASC", table)
	case "highest-price":
		queryText = fmt.Sprintf("SELECT * FROM %v Order By product_price DESC", table)
	case "a-z":
		queryText = fmt.Sprintf("SELECT * FROM %v Order By product_name ASC", table)
	case "z-a":
		queryText = fmt.Sprintf("SELECT * FROM %v Order By product_name DESC", table)
	default:
		queryText = fmt.Sprintf("SELECT * FROM %v Order By product_id DESC", table)
	}

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var product models.Product

		if err = rowQuery.Scan(&product.ProductId,
			&product.ProductName,
			&product.ProductPrice,
			&product.ProductDescription,
			&product.ProductQuantity); err != nil {
			return nil, err
		}

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil
}

// Insert Product
func Insert(ctx context.Context, prdct models.Product) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (product_name, product_price, product_description, product_quantity) values('%v',%v,'%v',%v)", table,
		prdct.ProductName,
		prdct.ProductPrice,
		prdct.ProductDescription,
		prdct.ProductQuantity)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Product
func Update(ctx context.Context, prdct models.Product) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set product_name = '%s', product_price = '%d', product_description = '%s', product_quantity = '%d' where product_id = %d",
		table,
		prdct.ProductName,
		prdct.ProductPrice,
		prdct.ProductDescription,
		prdct.ProductQuantity,
		prdct.ProductId,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Product
func Delete(ctx context.Context, prdct models.Product) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where product_id = '%d'", table, prdct.ProductId)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("product_id tidak ada")
	}

	return nil
}
