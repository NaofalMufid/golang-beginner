package controller

import (
	"database/sql"
	"mincha6/model"
	"time"
)

type ProductController struct {
	DB *sql.DB
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{DB: db}
}

func (pc *ProductController) AddProduct(product model.Product) error {
	queryInsert := `INSERT INTO products (name, created_at, updated_at) VALUES (?, ?, ?)`
	_, err := pc.DB.Exec(queryInsert, product.Name, time.Now(), time.Now())
	return err
}

func (pc *ProductController) UpdateProduct(product model.Product) error {
	queryUpdate := `UPDATE products SET name = ?, updated_at = ? WHERE id = ?`
	_, err := pc.DB.Exec(queryUpdate, product.Name, time.Now(), product.ID)
	return err
}

func (pc *ProductController) DeleteProduct(product model.Product) error {
	queryDelete := `DELETE FROM products WHERE id = ?`
	_, err := pc.DB.Exec(queryDelete, product.ID)
	return err
}

func (pc *ProductController) GetProducts() ([]model.Product, error) {
	querySelect := `SELECT * FROM products ORDER BY created_at ASC`
	rows, err := pc.DB.Query(querySelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (pc *ProductController) GetProductByID(id int) (model.Product, error) {
	querySelect := `SELECT * FROM products WHERE id = ?`
	var product model.Product
	err := pc.DB.QueryRow(querySelect, id).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (pc *ProductController) GetProductWithVariant() ([]model.Product, error) {
	querySelect := `
		SELECT p.id, p.name, v.id, v.variant_name, v.quantity
		FROM products p
		LEFT JOIN variants v ON p.id = v.product_id
		ORDER BY p.id ASC
	`
	rows, err := pc.DB.Query(querySelect)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	productVariant := make(map[int]model.Product)
	for rows.Next() {
		var productID int
		var name string

		var variantIDValid sql.NullInt64
		var variant_nameValid sql.NullString
		var quantityValid sql.NullInt64

		err := rows.Scan(&productID, &name, &variantIDValid, &variant_nameValid, &quantityValid)
		if err != nil {
			return nil, err
		}

		product, ok := productVariant[productID]
		if !ok {
			product = model.Product{
				ID:   productID,
				Name: name,
			}
		}

		if product.Variants == nil {
			product.Variants = make([]model.Variant, 0)
		}

		if variantIDValid.Valid {
			variant := model.Variant{
				ID:           int(variantIDValid.Int64),
				Variant_Name: variant_nameValid.String,
				Quantity:     int(quantityValid.Int64),
			}
			product.Variants = append(product.Variants, variant)
		}

		productVariant[productID] = product
	}

	products := make([]model.Product, 0, len(productVariant))
	for _, product := range productVariant {
		products = append(products, product)
	}

	return products, nil
}
