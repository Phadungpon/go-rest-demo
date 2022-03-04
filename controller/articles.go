package controller

import (
	"go-rest-api/models"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Articles struct {
}

type createarticles struct {
	Title string                `form:"title" binding:"required"`
	Body  string                `form:"body"  binding:"required"`
	Image *multipart.FileHeader `form:"image"  binding:"required"`
}

var articles []models.Article = []models.Article{
	{ID: 1, Title: "title#1", Body: "Body#1"},
	{ID: 2, Title: "title#2", Body: "Body#2"},
	{ID: 3, Title: "title#3", Body: "Body#3"},
	{ID: 4, Title: "title#4", Body: "Body#4"},
	{ID: 5, Title: "title#5", Body: "Body#5"},
}

func (a *Articles) FildAll(c *gin.Context) {
	result := articles
	if limit := c.Query("limit"); limit != "" {
		n, _ := strconv.Atoi(limit)
		result = result[:n]
	}

	c.JSON(http.StatusOK, gin.H{"articles": result})
}

func (a *Articles) FildOne(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, item := range articles {
		if item.ID == uint(id) {
			c.JSON(http.StatusOK, gin.H{"article": item})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
}

func (a *Articles) Create(c *gin.Context) {
	var form createarticles
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	article := models.Article{
		ID:    uint(len(articles) + 1),
		Title: form.Title,
		Body:  form.Body,
	}

	//get file
	file, _ := c.FormFile("image")
	//create file
	path := "uploads/articles/" + strconv.Itoa(int(article.ID))
	os.MkdirAll(path, 0755)
	//upload file
	filename := path + "/" + file.Filename
	err := c.SaveUploadedFile(file, filename)
	if err != nil {

	}
	//attach file to articles
	article.Image = os.Getenv("HOST") + "/" + filename

	articles = append(articles, article)
	c.JSON(http.StatusCreated, gin.H{"articles": article})
}