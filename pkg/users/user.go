package users

import (
	"github/SolBaa/gorm-api/models"

	"gorm.io/gorm"
)

//create a user
func CreateUser(db *gorm.DB, User *models.Cart) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetUsers(db *gorm.DB, User *[]models.Cart) (Carts *[]models.Cart, err error) {
	err = db.Find(User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}
