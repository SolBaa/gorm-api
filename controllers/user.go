package controllers

import (
	"encoding/json"
	"github/SolBaa/gorm-api/database"
	"github/SolBaa/gorm-api/models"
	"github/SolBaa/gorm-api/pkg/users"
	"github/SolBaa/gorm-api/viewmodels"
	"net/http"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Cart{}, &models.Product{})
	return &UserRepo{Db: db}
}

//create user
func (repository *UserRepo) CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.Cart
	json.NewDecoder(r.Body).Decode(&user)
	err := users.CreateUser(repository.Db, &user)
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(user)
}

//get users
func (repository *UserRepo) GetAllCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shoppingCart []models.Cart
	users, err := users.GetUsers(repository.Db, &shoppingCart)
	if err != nil {
		return
	}
	var cartsViewmodel []viewmodels.CartViewModel
	cart := viewmodels.CartViewModel{}
	for _, sc := range *users {
		cart = viewmodels.CartViewModel{
			CartID:     sc.ID,
			TotalPrice: sc.TotalPrice,
		}
		cartsViewmodel = append(cartsViewmodel, cart)
	}
	json.NewEncoder(w).Encode(cartsViewmodel)
}
