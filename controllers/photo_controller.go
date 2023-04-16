package controllers

import (
	"fmt"
	"myGram/database"
	"myGram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GetAllPhoto godoc
// @Summary Get all photos
// @Description Get all photos from given user Id
// @Tags photo
// @Accept json
// @Produce json
// @Security AuthToken
// @Success 200 {object} models.Photo
// @Router /photos [get]
func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	Photo := []models.Photo{}

	err := db.Preload("User").Where("user_id = ?", userId).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Currently this user didn't upload any photo yet!",
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// GetPhotoById godoc
// @Summary Get photo for a given Photo id
// @Description Get photo corresponding to the input Photo id
// @Tags photo
// @Accept json
// @Produce json
// @Security AuthToken
// @Param photoId path int true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photos/{photoId} [get]
func GetPhotoById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	age := int(userData["age"].(float64))

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}
	User := models.User{}

	err := db.Find(&Photo, photoId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "There is no photo with that id!",
		})
		return
	}

	db.First(&User, Photo.UserId)

	if Photo.UserId != userId {
		if User.Age >= 18 && age < 18 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "This photo can only be seen by 18+ years old!",
			})
			return
		}
	}

	c.JSON(http.StatusOK, Photo)
}

// CreatePhoto godoc
// @Summary Create photo from given data
// @Description Create new photo corresponding to the input data
// @Tags photo
// @Accept json
// @Produce json
// @Security AuthToken
// @Param models.Photo body models.Photo true "create photo"
// @Success 200 {object} models.Photo
// @Router /photos [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := c.ContentType()
	userId := uint(userData["id"].(float64))

	Photo := models.Photo{}
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	db.First(&User, userId)

	Photo.UserId = userId
	Photo.User = &User

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// UpdatePhoto godoc
// @Summary Update photo identified by the given Photo id
// @Description Update the photo corresponding to the input Photo id
// @Tags photo
// @Accept json
// @Produce json
// @Security AuthToken
// @Param photoId path int true "ID of the photo to be updated"
// @Success 200 {object} models.Photo
// @Router /photos/{photoId} [patch]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := c.ContentType()

	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserId = uint(userData["id"].(float64))
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// DeletePhoto godoc
// @Summary Delete photo identified by the given Photo id
// @Description Delete the photo corresponding to the input Photo id
// @Tags photo
// @Accept json
// @Produce json
// @Security AuthToken
// @Param photoId path int true "ID of the photo to be deleted"
// @Success 204 "No Content"
// @Router /photos/{photoId} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Comment := models.Comment{}

	err := db.Where("photoId=?", photoId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = db.Where("id=?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Photo with id %d has been successfully deleted", photoId),
	})
}
