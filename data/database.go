package data

import (
	"errors"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"neurothrone/go-db-web-api/models"
)

// Functions in PascalCase are exported and can be used by other packages.
// Functions in camelCase are private and can only be used within the package.

var db *gorm.DB

//func Init(file, server, database, username, password string, port int) {
//	if len(file) == 0 {
//		db = openMySql(server, database, username, password, port)
//	} else {
//		db, _ = gorm.Open(sqlite.Open(file), &gorm.Config{})
//	}
//
//	err := db.AutoMigrate(&models.Product{})
//	if err != nil {
//		fmt.Println("Error migrating Product")
//		return
//	}
//
//	generateSeedData()
//}

func Init(file, server, database, username, password string, port int) {
	fmt.Println("Initializing database with config:")
	fmt.Printf("File: %s, Server: %s, Database: %s, Username: %s, Port: %d\n",
		file, server, database, username, port)

	if len(file) == 0 {
		// Check if port is missing or zero
		if port == 0 {
			fmt.Println("Error: MySQL port is 0, defaulting to 3306")
			port = 3306 // Default MySQL port
		}

		db = openMySql(server, database, username, password, port)
	} else {
		fmt.Println("Using SQLite database file:", file)
		var err error
		db, err = gorm.Open(sqlite.Open(file), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("Failed to open SQLite database: %v", err))
		}
	}

	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		fmt.Println("Error migrating Product")
		return
	}

	generateSeedData()
}

func generateSeedData() {
	var productCount int64
	db.Model(&models.Product{}).Count(&productCount)

	if productCount != 0 {
		return
	}

	db.Create(&models.Product{Title: "Book", Price: 9.99, Description: "A good book", Category: "Books", Image: "book.jpg"})
	db.Create(&models.Product{Title: "Shirt", Price: 19.99, Description: "A good shirt", Category: "Clothes", Image: "shirt.jpg"})
	db.Create(&models.Product{Title: "Shoes", Price: 29.99, Description: "A good pair of shoes", Category: "Shoes", Image: "shoes.jpg"})
}

func openMySql(server, database, username, password string, port int) *gorm.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, server, port, database)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}

func CreateNewProduct(product *models.Product) *models.Product {
	db.Create(&product)
	return product
}

func GetAllProducts() []models.Product {
	var products []models.Product
	db.Find(&products)
	return products
}

func GetProduct(id int) *models.Product {
	var product models.Product
	err := db.First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &product
}

func UpdateProduct(product models.Product) bool {
	var dbProduct models.Product
	err := db.First(&dbProduct, product.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	dbProduct.Title = product.Title
	dbProduct.Price = product.Price
	dbProduct.Description = product.Description
	dbProduct.Category = product.Category
	dbProduct.Image = product.Image
	db.Save(&product)

	return true
}

func DeleteProduct(product *models.Product) {
	db.Delete(&product)
}
