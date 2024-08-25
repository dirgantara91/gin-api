package main

import (
	"database/sql"
	"fmt"

	_ "gin-api/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Regent struct {
	ID      int    `json:"idregent"`
	Name    string `json:"name"`
	Capital string `json:"capital"`
	Area    string `json:"area"`
}

var db *sql.DB

// @title Geography API
// @description API for managing geography data
// @version 1.0
// @host localhost:5000
// @BasePath /
func main() {
	// Connect to the database
	var err error
	db, err = sql.Open("mysql", "root:Kepandean@95@tcp(localhost:3306)/geography")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// Create a new Gin router
	r := gin.Default()

	// Register Swagger middleware
	// Register Swagger middleware
	//url := ginSwagger.URL("http://localhost:5000/swagger/doc.json") // url
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define the endpoints
	r.GET("/regent", getRegents)
	r.GET("/regent/:idregent", getRegent)

	// Run the API
	r.Run(":5000") // listen and serve on 0.0.0.0:5000
}

// GetRegents ... Get all Regents
// @Summary Get all Regents
// @Description get all Regents
// @Tags Regents
// @Success 200 {array} Regent
// @Failure 404 {object} object
// @Router /regent [get]
func getRegents(c *gin.Context) {
	rows, err := db.Query("SELECT idregent, name, capital, area FROM regent")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	var regents []Regent
	for rows.Next() {
		var u Regent
		err := rows.Scan(&u.ID, &u.Name, &u.Capital, &u.Area)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		regents = append(regents, u)
	}

	c.JSON(200, regents)
}

// getRegent ... Get the regent by idregent
// @Summary Get one regent
// @Description get regent by idregent
// @Tags Regents
// @Param idregent path string true "idregent"
// @Success 200 {object} Regent
// @Failure 400,404 {object} object
// @Router /regent/{idregent} [get]
func getRegent(c *gin.Context) {
	idregent := c.Param("idregent")
	row := db.QueryRow("SELECT idregent, name, capital, area FROM regent WHERE idregent = ?", idregent)

	var u Regent
	err := row.Scan(&u.ID, &u.Name, &u.Capital, &u.Area)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"error": "regent not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}
