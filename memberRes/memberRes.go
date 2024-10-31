package memberRes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reUnionBe/database"

	"github.com/gin-gonic/gin"
)

// User struct for user data
type MemberRes struct {
	NamaMember string `json:"namamember"`
	Skor1 int `json:"skor1"`
	Skor2 int `json:"skor2"`
	TotalSkor int `json:"totalskor"`
	Kategori string `json:"kategori"`
}

// addUser function to add a new user to the database
func AddMemberRes(ctx *gin.Context) {
	body := MemberRes{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Result is not defined"})
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Bad Input"})
		return
	}

	// Insert the user into the database
	_, err = database.Db.Exec("insert into scoringmemberres(namamember,skor1,skor2,totalskor,kategori) values ($1,$2,$3,$4,$5)", body.NamaMember, body.Skor1,body.Skor2,body.TotalSkor,body.Kategori)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Couldn't create the new user result."})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Result user is successfully created."})
	}
}

func GetUserRes(ctx *gin.Context){
	rows, err := database.Db.Query("SELECT namamember, skor1,skor2,totalskor,kategori FROM scoringmemberres")
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Could not retrieve result member"})
		return
	}
	defer rows.Close()

	var results []MemberRes //store data di sini

	for rows.Next() { //hasil dari query di simpan kedalam rows, trus di cek pake cursor
		var result MemberRes //ini untuk nyimpan data per ROW yang ditemuin di db
		if err := rows.Scan(&result.NamaMember, &result.Skor1, &result.Skor2, &result.TotalSkor, &result.Kategori); err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "Could not scan user result data"})
			return
		}
		results = append(results, result) //setiap data yg ketemu di row akan di append ke dalam array
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Error occurred while retrieving users results"})
		return
	}

	ctx.JSON(http.StatusOK, results)
}