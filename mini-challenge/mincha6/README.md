# Mini Challenge 6
- Membuat CRUD Product dan Variant hanya dengan driver mysql tanpa orm misal ini [mysql driver](github.com/go-sql-driver/mysql)
- program berjalan cukup dengan comment / uncoment function saja

# Cara Menjalankan program
- Program Mengguanakan database **mysql**
- Clone repository ini masuk ke folder mini-challenge->mincha6
- sesuaikan isi file env seperti db user, db password dan db name
- database yang sudah jadi ada di folder database
- menggunakan godotenv bisa di install ```go get github.com/joho/godotenv```
## command menjalankan program sederhana ini
1. Product
    - Get All Product ```  go run main.go all-product ```
    - Get Product By ID ```  go run main.go get-product [id] ```
    - Create Product ```  go run main.go new-product "product name" ```
    - Edit Product ``` go run main.go edit-product [id] "name" ```
    - Delete Product ``` go run main.go delete-product [id] ```
2. Variant
    - Product With Variant ```  go run main.go prodict-variant ```
    - Create Variant ```  go run main.go new-variant "product name" [quantity] [product_id] ```
    - Edit Variant ``` go run main.go edit-variant [id] "variant name" [quantity] [product_id] ```
    - Delete Variant ``` go run main.go delete-variant [id] ```