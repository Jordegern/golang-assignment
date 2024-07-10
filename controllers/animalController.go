package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"uniwise/animals/initializers"
	"uniwise/animals/models"
)

func CreateAnimal(ctx *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var animal models.Animal
	if err := ctx.ShouldBindJSON(&animal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(animal)
	result := initializers.DB.Create(&animal) // pass pointer of data to Create

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating animal"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Animal created successfully", "animal": animal})

}
func ListAnimals(ctx *gin.Context) {
	var animals []models.Animal
	initializers.DB.Find(&animals)
	ctx.JSON(http.StatusOK, gin.H{"message": "All animals", "animals": animals})
}
func GetAnimal(ctx *gin.Context) {
	var animal models.Animal
	id := ctx.Param("id")

	initializers.DB.First(&animal, "id = ?", id)
	ctx.JSON(http.StatusOK, gin.H{"message": "Animal found", "animal": animal})
}
func UpdateAnimal(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		Birthday string `json:"birthday"`
		Gender   string `json:"gender"`
		Status   string `json:"status"`
		Race     string `json:"race"`
		Weight   string `json:"weight"`
	}
		ctx.Bind(&input)

		var animal models.Animal
		if err:= initializers.DB.First(&animal, id).Error; err == nil {
		    ctx.JSON(http.StatusNotFound, gin.H{"message": "Animal not found"})
		    return
	    }
		if err := initializers.DB.Model(&animal).Updates(&models.Animal{Birthday: input.Birthday, Gender: input.Gender, Status: input.Status, Race: input.Race, Weight: input.Weight}).Error; err != nil {
		    ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update animal"})
		    return
	    }
		ctx.JSON(http.StatusOK, gin.H{"message": "Animal updated", "animal": animal})
	}
}
func DeleteAnimal(ctx *gin.Context) {
	id := ctx.Param("id")

	if(initializers.DB.First(&models.Animal{}, id).RowsAffected == 0){
		ctx.JSON(404, gin.H{"message": "Animal not found"})
		return
	}

	initializers.DB.Delete(&models.Animal{}, id)

	ctx.JSON(200, gin.H{"message": "Animal deleted successfully"})
}
