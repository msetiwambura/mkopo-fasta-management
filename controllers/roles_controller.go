package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "time"
	"usrmanagement/configs"
	"usrmanagement/models"
	"usrmanagement/utils"
)

func CreateRole(c *gin.Context) {
	var newRole models.Role
	if err := c.ShouldBindJSON(&newRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := configs.DB.Create(&newRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	response := utils.CreateSuccessResponse("Success", "Request processed successfully", []models.Role{newRole})

	c.JSON(http.StatusCreated, response)
}

func GetRoles(c *gin.Context) {
	var roles []models.Role
	configs.DB.Preload("Permissions").Preload("Pages").Find(&roles)

	response := utils.CreateSuccessResponse("Success", "Request processed successfully", roles)
	c.JSON(http.StatusOK, response)
}
