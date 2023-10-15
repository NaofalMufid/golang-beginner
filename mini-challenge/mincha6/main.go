package main

import (
	"fmt"
	"mincha6/controller"
	"mincha6/database"
	"mincha6/model"
	"os"
	"strconv"
	"strings"
)

func main() {
	db, err := database.StartDB()
	if err != nil {
		fmt.Println("Error intializing database:", err)
		panic(err)
	}

	defer db.Close()

	productController := controller.NewProductController(db)
	variantController := controller.NewVarianController(db)

	// call function controller with cli argument
	if len(os.Args) < 2 {
		fmt.Println("Penggunaan: go run main.go [command]")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	// PRODUCT COMMAND
	case "all-product":
		if len(os.Args) != 2 {
			fmt.Println("Command untuk all product kurang tepat")
			fmt.Println("Contoh: go run main.go all-product")
			os.Exit(1)
		}

		data, err := productController.GetProducts()
		if err != nil {
			fmt.Println("Error get all product", err)
		}

		fmt.Println(strings.Repeat("=", 10))
		for _, v := range data {
			fmt.Println("ID:", v.ID)
			fmt.Println("Name:", v.Name)
			fmt.Println("CreatedAt:", v.CreatedAt)
			fmt.Println("UpdatedAt:", v.UpdatedAt)
			fmt.Println(strings.Repeat("=", 10))
		}
	case "get-product":
		if len(os.Args) != 3 {
			fmt.Println("Command untuk get product kurang tepat")
			fmt.Println("Contoh: go run main.go get-product [id]")
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error argument param id", err)
		}

		data, err := productController.GetProductByID(id)
		if err != nil {
			fmt.Println("Error get product by id", err)
		}

		fmt.Println(strings.Repeat("=", 10))
		fmt.Println("ID: ", data.ID)
		fmt.Println("Name: ", data.Name)
		fmt.Println("CreatedAt: ", data.CreatedAt)
		fmt.Println("UpdatedAt: ", data.UpdatedAt)
		fmt.Println(strings.Repeat("=", 10))
	case "new-product":
		if len(os.Args) != 3 {
			fmt.Println("Command untuk new product kurang tepat")
			fmt.Println(`Penggunaan: go run main.go new-product "name"`)
			os.Exit(1)
		}

		name := os.Args[2]

		product := model.Product{
			Name: name,
		}

		err := productController.AddProduct(product)
		if err != nil {
			fmt.Println("Error create new product", err)
		} else {
			fmt.Println("Product created successfully")
		}
	case "edit-product":
		if len(os.Args) != 4 {
			fmt.Println("Command untuk edit product kurang tepat")
			fmt.Println(`Contoh: go run main.go edit-product [id] "name"`)
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error argument param id", err)
		}
		name := os.Args[3]

		productToUpdate, err := productController.GetProductByID(id)
		if err != nil {
			fmt.Println("Gagal mendapatkan product:", err)
			return
		}

		productToUpdate.Name = name

		err = productController.UpdateProduct(productToUpdate)
		if err != nil {
			fmt.Println("Error updating product", err)
		} else {
			fmt.Println("Product updated successfully")
		}
	case "delete-product":
		if len(os.Args) != 3 {
			fmt.Println("Command untuk delete product kurang tepat")
			fmt.Println("Contoh: go run main.go delete-product [id]")
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error argument param id", err)
		}

		product, err := productController.GetProductByID(id)
		if err != nil {
			fmt.Println("Gagal mendapatkan product:", err)
			return
		}

		err = productController.DeleteProduct(product)
		if err != nil {
			fmt.Println("Error deleting product", err)
		} else {
			fmt.Println("Product deleted successfully")
		}
	// VARIANT COMMAND
	case "product-variant":
		if len(os.Args) != 2 {
			fmt.Println("Command untuk product-variant kurang tepat")
			fmt.Println("Contoh: go run main.go prodict-variant")
			os.Exit(1)
		}

		data, err := productController.GetProductWithVariant()
		if err != nil {
			fmt.Println("Error product with variant", err)
		}

		fmt.Println(strings.Repeat("=", 10))
		for _, v := range data {
			fmt.Println("Product ID :", v.ID)
			fmt.Println("Product Name :", v.Name)
			fmt.Println("Variants :", v.Variants)
			fmt.Println(strings.Repeat("=", 10))
		}
	case "new-variant":
		if len(os.Args) != 5 {
			fmt.Println("Command untuk new variant kurang tepat")
			fmt.Println(`Penggunaan: go run main.go new-variant "variant name" [quantity] [product id]`)
			os.Exit(1)
		}

		variant_name := os.Args[2]
		quantity, _ := strconv.Atoi(os.Args[3])
		product_id, _ := strconv.Atoi(os.Args[4])

		variant := model.Variant{
			Variant_Name: variant_name,
			Quantity:     quantity,
			Product_ID:   product_id,
		}

		err := variantController.Addvariant(variant)
		if err != nil {
			fmt.Println("Error create new variant", err)
		} else {
			fmt.Println("Variant created successfully")
		}
	case "edit-variant":
		if len(os.Args) != 6 {
			fmt.Println("Command untuk edit variant kurang tepat")
			fmt.Println(`Contoh: go run main.go edit-variant [id] "variant name" [quantity] [product_id]`)
			os.Exit(1)
		}

		id, _ := strconv.Atoi(os.Args[2])
		variant_name := os.Args[3]
		quantity, _ := strconv.Atoi(os.Args[4])
		product_id, _ := strconv.Atoi(os.Args[5])

		variantToUpdate, err := variantController.GetvariantByID(id)
		if err != nil {
			fmt.Println("Gagal mendapatkan variant:", err)
			return
		}

		variantToUpdate.Variant_Name = variant_name
		variantToUpdate.Quantity = quantity
		variantToUpdate.Product_ID = product_id

		err = variantController.Updatevariant(variantToUpdate)
		if err != nil {
			fmt.Println("Error updating variant", err)
		} else {
			fmt.Println("Variant updated successfully")
		}
	case "delete-variant":
		if len(os.Args) != 3 {
			fmt.Println("Command untuk delete product kurang tepat")
			fmt.Println("Contoh: go run main.go delete-product [id]")
			os.Exit(1)
		}

		id, _ := strconv.Atoi(os.Args[2])

		variant, err := variantController.GetvariantByID(id)
		if err != nil {
			fmt.Println("Gagal mendapatkan variant:", err)
			return
		}

		err = variantController.Deletevariant(variant)
		if err != nil {
			fmt.Println("Error deleting variant", err)
		} else {
			fmt.Println("Variant deleted successfully")
		}
	default:
		fmt.Println("Command tidak sesuai:", command)
	}
}
