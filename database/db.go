package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
)

var Db *sql.DB 


func ConnectDatabase() {

   err := godotenv.Load()//load .env file
   if err != nil {
      fmt.Println("Error is occurred  on .env file please check")
   }
   //we read our .env file
   host := os.Getenv("HOST")
   port, _ := strconv.Atoi(os.Getenv("PORT")) // untuk PORT harus di convert jadi INT terlebih dahulu pake Atoi()
   user := os.Getenv("USER")
   dbname := os.Getenv("DB_NAME")
   pass := os.Getenv("PASSWORD")

   // set up postgres sql to open it.
   psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
       host, port, user, dbname, pass)
   db, errSql := sql.Open("postgres", psqlSetup)
   if errSql != nil {
      fmt.Println("Error connecting database ", err)
      panic(err)
   } else {
      Db = db
      fmt.Println("Connected")
   }
}