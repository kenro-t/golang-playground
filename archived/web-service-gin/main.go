package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"fmt"
)

type album struct {
	ID string `json:"id`
	Title string `json:"title`
	Artist string `json:"artist`
	Price float64 `json:"price`
}

type test struct {
	test string 
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

	for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
    var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// フィールド名は大文字で始めないとdecode encodeされない
type JsonRequest struct {
	Test  string `json:"test"`
}
func postJsonTest(c *gin.Context) {

	var postJson JsonRequest
	if err := c.ShouldBindJSON(&postJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"test": postJson.Test})
    // c.IndentedJSON(http.StatusOK, gin.H{"message": message})
}

func postFormTest(c *gin.Context) {
    var newTest test

	
    if err := c.ShouldBind(&newTest); err != nil {
        // return
    }
	s := c.PostForm("test")
	message := fmt.Sprintf("%v", s)
	
    c.IndentedJSON(http.StatusOK, gin.H{"message": message})
}


func main() {
    router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

    router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
	router.POST("/post-form-test", postFormTest)
	router.POST("/post-json-test", postJsonTest)
    router.Run("localhost:8080")
}

/*
備忘録

対象パスの依存関係を解消する
go get パス

jsonで値を受け取る時、structのフィールド名は大文字で始めないとdecode encodeされない
*/