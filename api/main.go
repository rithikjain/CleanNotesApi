package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/rithikjain/CleanNotesApi/api/handler"
	"github.com/rithikjain/CleanNotesApi/pkg/user"
	"log"
	"net/http"
)

func dbConnect(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode),
	)

	return db, err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := dbConnect("localhost", "5432", "postgres", "notesapi", "30june", "disable")
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
	defer db.Close()
	fmt.Println("Connected to DB...")

	db.LogMode(true)
	db.AutoMigrate(&user.User{})

	// initializing repos and services
	userRepo := user.NewRepo(db)
	userSvc := user.NewService(userRepo)

	r := http.NewServeMux()

	handler.MakeUserHandler(r, userSvc)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	})

	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
