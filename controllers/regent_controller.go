package controller

import (
	_ "gin-api/docs"
	model "gin-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RegentController struct{}

// GetRegents ... Get all Regents
// @Summary Get all Regents
// @Description get all Regents
// @Tags Regents
// @Success 200 {array} Regent
// @Failure 404 {object} object
// @Router /regent [get]
func (rc *RegentController) GetRegents(c *gin.Context) {
	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{
			"errors": []gin.H{
				{"title": "Internal Server Error", "detail": err.Error()},
			},
		})
		return
	}
	defer db.Close()

	regents, err := model.Regent{}.GetRegents(db)
	if err != nil {
		c.JSON(500, gin.H{
			"errors": []gin.H{
				{"title": "Internal Server Error", "detail": err.Error()},
			},
		})
		return
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
func (rc *RegentController) GetRegent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer db.Close()

	regent := &model.Regent{}
	result, err := regent.GetRegent(db, id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if result == nil {
		c.JSON(404, gin.H{
			"error": "Regent not found",
		})
		return
	}

	c.JSON(200, result)
}
