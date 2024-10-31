package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reUnionBe/database"

	"github.com/gin-gonic/gin"
)

// User struct for user data
type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// addUser function to add a new user to the database
func AddUser(ctx *gin.Context) {
	body := User{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "User is not defined"})
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Bad Input"})
		return
	}

	// Insert the user into the database
	_, err = database.Db.Exec("insert into users(email,password) values ($1,$2)", body.Email, body.Password)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Couldn't create the new user."})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User is successfully created."})
	}
}

func GetUser(ctx *gin.Context){
	rows, err := database.Db.Query("SELECT email, password FROM users")
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Could not retrieve users"})
		return
	}
	defer rows.Close()

	var users []User //store data di sini

	for rows.Next() { //hasil dari query di simpan kedalam rows, trus di cek pake cursor
		var user User //ini untuk nyimpan data per ROW yang ditemuin di db
		if err := rows.Scan(&user.Email, &user.Password); err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "Could not scan user data"})
			return
		}
		users = append(users, user) //setiap data yg ketemu di row akan di append ke dalam array
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Error occurred while retrieving users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}