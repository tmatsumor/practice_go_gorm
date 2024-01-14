package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"uniqueIndex"`
}

func main() {
	dsn := "user=postgres password=postgres host=localhost dbname=test sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// テーブルのマイグレーション
	db.AutoMigrate(&User{})

	// GORMを使ったデータベース操作をここに記述
	/*
		// create user
		user := User{Name: "John Doe", Email: "joen@example.com"}
		result := db.Create(&user)
		if result.Error != nil {
			panic("failed to insert record")
		}
	*/
	// select user
	var user User
	db.First(&user)
	db.Where("name=?", "John Doe").Find(&user)
	var users []User
	db.Find(&users)

	// update user
	user.Name = "Jane Doe"
	db.Save(&user)

	db.Model(&user).Update("email", "hoge@example.com")

	db.Model(&user).Updates(map[string]interface{}{"name": "Jane DDD", "email": "foo@example.com"})

	// delete user
	db.Where("email=? AND name=?", "foo@example.com", "Jane DDD").Delete(&User{})
}
