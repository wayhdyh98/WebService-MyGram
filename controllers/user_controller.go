package controllers

import (
	"myGram/database"
	"myGram/helpers"
	"myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func HasData(u *models.User) bool {
	if u.ID != 0 || u.Username != "" || u.Email != "" {
		return true
	}
	return false
}

// UserRegister godoc
// @Summary Create users from given data
// @Description Create new user corresponding to the input data
// @Tags users
// @Accept json
// @Produce json
// @Param models.User body models.User true "create user"
// @Success 200 {object} models.User
// @Router /users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := c.ContentType()

	TempUser, User := models.User{}, models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	db.Where("email = ?", User.Email).First(&TempUser)
	checkUser := HasData(&TempUser)

	if checkUser {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Email is already registered!",
		})
		return
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age": User.Age,
	})
}

// UserLogin godoc
// @Summary Get users for a given Id
// @Description Get user profile corresponding to the input Id
// @Tags users
// @Accept json
// @Produce json
// @Param models.User body models.User true "login user"
// @Success 200 {object} models.User
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := c.ContentType()

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password := User.Password

	err := db.Where("email = ?", User.Email).First(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Email is not found!",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Password incorrect!",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.Age)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
