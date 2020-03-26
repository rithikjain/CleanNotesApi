package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rithikjain/CleanNotesApi/api/handler"
	"github.com/rithikjain/CleanNotesApi/pkg/note"
	"github.com/rithikjain/CleanNotesApi/pkg/user"
	"log"
	"net/http"
	"os"
)

func dbConnect(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {

	// In the case of heroku
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	}
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode),
	)

	return db, err
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("INFO: No PORT environment variable detected, defaulting to 3000")
		return "localhost:3000"
	}
	return ":" + port
}

func main() {
	/*
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	*/

	db, err := dbConnect(
		os.Getenv("dbHost"),
		os.Getenv("dbPort"),
		os.Getenv("dbUser"),
		os.Getenv("dbName"),
		os.Getenv("dbPass"),
		os.Getenv("sslmode"),
	)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
	defer db.Close()
	fmt.Println("Connected to DB...")

	db.LogMode(true)
	db.AutoMigrate(&user.User{}, &note.Note{})

	// initializing repos and services
	userRepo := user.NewRepo(db)
	noteRepo := note.NewRepo(db)

	userSvc := user.NewService(userRepo)
	noteSvc := note.NewService(noteRepo)

	r := http.NewServeMux()

	handler.MakeUserHandler(r, userSvc)
	handler.MakeNoteHandler(r, noteSvc)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello There"))
		return
	})

	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(GetPort(), r))
}
