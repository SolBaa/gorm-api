package main

import (
	"github/SolBaa/gorm-api/controllers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")

	// dsn := fmt.Sprintf(fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword))
	// dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("Cannot connect to database")
	// 	log.Fatal("This is the error:", err)
	// } else {
	// 	fmt.Println("We are connected to the database")
	// }

	userRepo := controllers.New()
	r := mux.NewRouter()
	r.HandleFunc("/cart", userRepo.CreateCart).Methods("POST")
	r.HandleFunc("/cart", userRepo.GetAllCarts).Methods("GET")
	// r.HandleFunc("/user").Methods("GET")

	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
