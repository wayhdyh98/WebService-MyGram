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

// GetAllComment godoc
// @Summary Get all comments
// @Description Get all comments from given user Id
// @Tags comment
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comments [get]
func GetAllComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	Comment := []models.Comment{}

	err := db.Preload("User").Where("user_id = ?", userId).Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "No comment from this user yet!",
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// GetCommentById godoc
// @Summary Get comment for a given Comment id
// @Description Get comment corresponding to the input Comment id
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment"
// @Success 200 {object} models.Comment
// @Router /comments/{commentId} [get]
func GetCommentById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	err := db.Find(&Comment, commentId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "No comment for that id!",
		})
		return
	}

	if Comment.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Can't access others comment!",
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// CreateComment godoc
// @Summary Create comment from given data
// @Description Create new comment corresponding to the input data
// @Tags comment
// @Accept json
// @Produce json
// @Param models.Comment body models.Comment true "create comment"
// @Success 200 {object} models.Comment
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := c.ContentType()
	userId := uint(userData["id"].(float64))

	Comment := models.Comment{}
	Photo := models.Photo{}
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	db.First(&User, userId)
	db.First(&Photo, Comment.PhotoId)

	Comment.UserId = userId
	Comment.User = &User
	Comment.Photo = &Photo

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// UpdateComment godoc
// @Summary Update comment identified by the given Comment id
// @Description Update the comment corresponding to the input Comment id
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment to be updated"
// @Success 200 {object} models.Comment
// @Router /comments/{commentId} [patch]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := c.ContentType()

	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserId = uint(userData["id"].(float64))
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// DeleteComment godoc
// @Summary Delete comment identified by the given Comment id
// @Description Delete the comment corresponding to the input Comment id
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment to be deleted"
// @Success 204 "No Content"
// @Router /comments/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.Where("id=?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("This comment with id %d has been deleted", commentId),
	})
}
