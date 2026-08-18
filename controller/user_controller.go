package controller

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"strconv"

	"ruibao_pos_system/models"
	"ruibao_pos_system/service"
	"ruibao_pos_system/utils"
)

func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := service.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetUserByID(c *gin.Context) {
	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	userID := uint(parsedID)

	user, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Admin    bool   `json:"admin"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	user, err := service.CreateUser(input.Username, input.Password, input.Admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Admin    bool   `json:"admin"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to secure password"})
		return
	}

	user := models.User{
		ID:       uint(parsedID),
		Username: input.Username,
		Password: hashedPassword,
		Admin:    input.Admin,
	}

	err = service.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = service.DeleteUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}