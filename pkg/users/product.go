package users

import (
	"github/SolBaa/gorm-api/models"

	"gorm.io/gorm"
)

//create a user
func CreateProduct(db *gorm.DB, User *models.Product) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetProduct(db *gorm.DB, User *[]models.Product) (Carts *[]models.Product, err error) {
	err = db.Find(User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}
