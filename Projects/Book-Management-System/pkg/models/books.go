package models

import (
	"log"

	"github.com/devilzs1/book-store/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("Id=?", Id).Find(&book)
	return &book, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("Id=?", Id).Delete(&book)
	return book
}

func UpdateBook(b *Book, updates map[string]interface{}) (*Book) {
    var existingBook Book
    if err := db.First(&existingBook, b.ID).Error; err != nil {
		log.Fatal(err)
        return nil
    }
    if err := db.Model(&existingBook).Updates(updates).Error; err != nil {
		log.Fatal(err)
        return nil
    }
    return &existingBook
}

