package controller

import (
	"database/sql"
	"mincha6/model"
	"time"
)

type VarianController struct {
	DB *sql.DB
}

func NewVarianController(db *sql.DB) *VarianController {
	return &VarianController{DB: db}
}

func (pc *VarianController) Addvariant(variant model.Variant) error {
	queryInsert := `INSERT INTO variants (variant_name, quantity, product_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	_, err := pc.DB.Exec(queryInsert, variant.Variant_Name, variant.Quantity, variant.Product_ID, time.Now(), time.Now())
	return err
}

func (pc *VarianController) Updatevariant(variant model.Variant) error {
	queryUpdate := `UPDATE variants SET variant_name = ?, quantity = ?, product_id = ?, updated_at = ? WHERE id = ?`
	_, err := pc.DB.Exec(queryUpdate, variant.Variant_Name, variant.Quantity, variant.Product_ID, time.Now(), variant.ID)
	return err
}

func (pc *VarianController) Deletevariant(variant model.Variant) error {
	queryDelete := `DELETE FROM variants WHERE id = ?`
	_, err := pc.DB.Exec(queryDelete, variant.ID)
	return err
}

func (pc *VarianController) GetvariantByID(id int) (model.Variant, error) {
	querySelect := `SELECT * FROM variants WHERE id = ?`
	var variant model.Variant
	err := pc.DB.QueryRow(querySelect, id).Scan(&variant.ID, &variant.Variant_Name, &variant.Quantity, &variant.Product_ID, &variant.CreatedAt, &variant.UpdatedAt)
	if err != nil {
		return model.Variant{}, err
	}

	return variant, nil
}
