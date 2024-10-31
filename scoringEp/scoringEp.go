package scoringEp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reUnionBe/database"

	"github.com/gin-gonic/gin"
)

// User struct for user data
type ScoringEp struct {
	Title string `json:"title"`
	Date string `json:"date"`
	Location string `json:"location"`
	Time string `json:"time"`
	ScoringType string `json:"scoringtype"`
}

// addUser function to add a new user to the database
func AddScoringEp(ctx *gin.Context) {
	body := ScoringEp{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Scoring Ep is not defined"})
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Bad Input"})
		return
	}

	// Insert the user into the database
	_, err = database.Db.Exec("insert into scoringep(title,date,location,time,scoringtype) values ($1,$2,$3,$4,$5)", body.Title, body.Date,body.Location,body.Time,body.ScoringType)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Couldn't create the new Scoring Episode."})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Scoring Ep is successfully created."})
	}
}

func GetScoringEp(ctx *gin.Context){
	rows, err := database.Db.Query("SELECT title, date,location,time,scoringtype FROM scoringep")
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Could not retrieve Scoring Episode"})
		return
	}
	defer rows.Close()

	var results []ScoringEp //store data di sini

	for rows.Next() { //hasil dari query di simpan kedalam rows, trus di cek pake cursor
		var result ScoringEp //ini untuk nyimpan data per ROW yang ditemuin di db
		if err := rows.Scan(&result.Title, &result.Date, &result.Location, &result.Time, &result.ScoringType); err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "Could not scan Scoring Episode data"})
			return
		}
		results = append(results, result) //setiap data yg ketemu di row akan di append ke dalam array
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Error occurred while retrieving Scoring Episode "})
		return
	}

	ctx.JSON(http.StatusOK, results)
}