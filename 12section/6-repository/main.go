package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"go-course.com/12section/6-repository/repository"
)

func main() {

	dbName := "users_database.db"
	db, err := connectToDatabase(dbName)
	checkErr(err)

	defer func() {
		fmt.Println("Closing database")
		if err := db.Close(); err != nil {
			log.Printf("error closing database connection: %v", err)
		}
	}()

	repo := repository.NewSQLUserRepository(db)

	user, err := repo.CreateUser("User from repo2", "userrepo2@gmail.com", "asdkdfsja", "avatar123")
	checkErr(err)

	fmt.Printf("%v\n", user)

	printUsers(repo)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func printUsers(repo repository.UserRepository) {
	users, err := repo.GetUsers()

	checkErr(err)
	for _, user := range users {
		fmt.Println(user.ID, user.Email)
	}
}
