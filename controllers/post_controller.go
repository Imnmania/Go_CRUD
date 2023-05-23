package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnmania/go_crud/initializers"
	"github.com/imnmania/go_crud/models"
)

func PostsCreate(ctx *gin.Context) {
	// Get data off request body
	var body struct {
		Title string
		Body  string
	}
	ctx.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}

	// Check if post exists by title
	initializers.DB.First(&post, "title = ?", body.Title)
	if post.ID != 0 {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"status":  "Failed",
			"message": "Post with this title already exists!",
		})
		return
	}

	// Create a post
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Post creation failed!",
		})
		return
	}

	// Send it back
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "Success",
		"data":   post,
	})
}

func PostsGetAll(ctx *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Send response
	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"data":   posts,
	})
}

func PostsGetById(ctx *gin.Context) {
	// Get the id from path param
	id := ctx.Param("id")

	// Get from DB
	var post models.Post
	initializers.DB.First(&post, id)
	if post.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Failed",
			"message": "Post does not exist!",
		})
		return
	}

	// Return with response
	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"data":   post,
	})
}

func PostsUpdate(ctx *gin.Context) {
	// Get the id off the path params
	id := ctx.Param("id")

	// Get the data off the req body
	var body struct {
		Title string
		Body  string
	}
	ctx.Bind(&body)

	// Find the post
	var post models.Post
	initializers.DB.First(&post, id)
	if post.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Post does not exist!",
		})
		return
	}

	// Update the post
	result := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Failed to update post!",
		})
		return
	}

	// Send the response
	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"data":   post,
	})
}

func PostsDelete(ctx *gin.Context) {
	// Get the id off path params
	id := ctx.Param("id")

	// Find the post if exist
	var post models.Post
	initializers.DB.First(&post, id)
	if post.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Failed",
			"message": "Post does not exist!",
		})
		return
	}

	// Delete from DB
	result := initializers.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": result.Error,
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Post deleted",
	})

}
