package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"usrmanagement/configs"
	"usrmanagement/models"
	"usrmanagement/utils"
)

func CreatePage(c *gin.Context) {
	var newPage models.Page
	if err := c.ShouldBindJSON(&newPage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := configs.DB.Create(&newPage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create page"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Page created successfully", "page": newPage})
}

func GetUserPages(c *gin.Context) {
	claims := c.MustGet("claims").(*utils.Claims)

	var role models.Role
	if err := configs.DB.Preload("Pages").First(&role, claims.Role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching pages"})
		return
	}
	res := utils.CreateSuccessResponse("Success", "Request Processed Successfully", role.Pages)

	c.JSON(http.StatusOK, res)
}

func AssignPagesToRole(c *gin.Context) {
	var input struct {
		RoleID  uint   `json:"roleId"`
		PageIDs []uint `json:"pageIds"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var role models.Role
	if err := configs.DB.Preload("Pages").Where("id = ?", input.RoleID).First(&role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	var pages []models.Page
	if err := configs.DB.Where("id IN ?", input.PageIDs).Find(&pages).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Some pages not found"})
		return
	}

	role.Pages = pages
	if err := configs.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign pages to role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pages assigned to role successfully"})
}
