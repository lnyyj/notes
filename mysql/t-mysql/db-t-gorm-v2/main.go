package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User x
type User struct {
	Name string `gorm:"column:name;"`
	Age  int    `gorm:"column:age;"`
	Role int    `gorm:"column:role;"`

	gorm.Model
}

func main() {
	db, err := gorm.Open(mysql.Open("root:1234qwer@tcp(localhost:3306)/admin"), &gorm.Config{})
	if err != nil {
		fmt.Println("------>err: ", err)
		return
	}

	// 大部分 CRUD API 都是兼容的
	db.AutoMigrate(&User{})

	// user := Product{Name: "ssy"}
	// db.Create(&user)
	// db.First(&user, 1)
	// db.Model(&user).Update("Age", 18)
	// db.Model(&user).Omit("Role").Updates(map[string]interface{}{"Name": "jinzhu", "Role": "admin"})
	// db.Delete(&user)
}
