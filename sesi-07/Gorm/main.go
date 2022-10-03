package main

import (
	"Gorm/database"
	"Gorm/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createUser("johndoe@gmail.com")
	// getUserById(1)
	// updateUserById("johnjohn@gmail.com")
	// createProduct(1, "YLO", "YYY")
	// getUsersWithProduct()
	deleteProductById(1)
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data : ", err)
		return
	}

	fmt.Println("New User Data : ", User)
}

func getUserById(id uint) {
	db := database.GetDB()
	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user: ", err)
	}
	fmt.Printf("User Data : %+v \n", user)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()
	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data", err)
		return
	}
	fmt.Printf("Update user's email : %+v \n", user.Email)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err)
		return
	}

	fmt.Println("New Product Data:", Product)
}

func getUsersWithProduct() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with product", err)
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error Deleting Product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted\n\n", id)
}
