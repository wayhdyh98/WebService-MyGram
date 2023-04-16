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

// GetAllMedia godoc
// @Summary Get all medias
// @Description Get all medias from given user Id
// @Tags media
// @Accept json
// @Produce json
// @Success 200 {object} models.Socialmedia
// @Router /medias [get]
// @Security Bearer
func GetAllMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	Media := []models.Socialmedia{}

	err := db.Preload("User").Where("user_id = ?", userId).Find(&Media).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "No social media has been found yet from this user!",
		})
		return
	}

	c.JSON(http.StatusOK, Media)
}

// GetMediaById godoc
// @Summary Get media for a given Media id
// @Description Get media corresponding to the input Media id
// @Tags media
// @Accept json
// @Produce json
// @Param mediaId path int true "ID of the media"
// @Success 200 {object} models.Socialmedia
// @Router /medias/{mediaId} [get]
// @Security Bearer
func GetMediaById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	mediaId, _ := strconv.Atoi(c.Param("mediaId"))
	Media := models.Socialmedia{}

	err := db.Find(&Media, mediaId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "No Social media for that id!",
		})
		return
	}

	if Media.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Can't access others Social media!",
		})
		return
	}

	c.JSON(http.StatusOK, Media)
}

// CreateMedia godoc
// @Summary Create media from given data
// @Description Create new media corresponding to the input data
// @Tags media
// @Accept json
// @Produce json
// @Param models.Socialmedia body models.Socialmedia true "create media"
// @Success 200 {object} models.Socialmedia
// @Router /medias [post]
// @Security Bearer
func CreateMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := c.ContentType()
	userId := uint(userData["id"].(float64))

	Media := models.Socialmedia{}
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Media)
	} else {
		c.ShouldBind(&Media)
	}

	db.First(&User, userId)

	Media.UserId = userId
	Media.User = &User

	err := db.Debug().Create(&Media).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Media)
}

// UpdateMedia godoc
// @Summary Update media identified by the given Media id
// @Description Update the media corresponding to the input Media id
// @Tags media
// @Accept json
// @Produce json
// @Param mediaId path int true "ID of the media to be updated"
// @Success 200 {object} models.Socialmedia
// @Router /medias/{mediaId} [patch]
// @Security Bearer
func UpdateMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := c.ContentType()

	Media := models.Socialmedia{}
	mediaId, _ := strconv.Atoi(c.Param("mediaId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Media)
	} else {
		c.ShouldBind(&Media)
	}

	Media.UserId = uint(userData["id"].(float64))
	Media.ID = uint(mediaId)

	err := db.Model(&Media).Updates(models.Socialmedia{Name: Media.Name, SocialMediaUrl: Media.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Media)
}

// DeleteMedia godoc
// @Summary Delete media identified by the given Media id
// @Description Delete the media corresponding to the input Media id
// @Tags media
// @Accept json
// @Produce json
// @Param mediaId path int true "ID of the media to be deleted"
// @Success 204 "No Content"
// @Router /medias/{mediaId} [delete]
// @Security Bearer
func DeleteMedia(c *gin.Context) {
	db := database.GetDB()
	Media := models.Socialmedia{}

	mediaId, _ := strconv.Atoi(c.Param("mediaId"))

	err := db.Where("id=?", mediaId).Delete(&Media).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("This Media with id %d has been deleted", mediaId),
	})
}
