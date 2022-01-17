package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri))
	if err != nil {
		fmt.Print(err)
	}

	DB = conn
	DB.Debug().AutoMigrate(&Account{}, &Travel{}, &Election{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return DB
}
