package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type aboutapi struct {
	App     string  `json:"app"`
	Github  string  `json:"github"`
	Version float64 `json:"version"`
}

var aboutapp = []aboutapi{
	{App: "SuperRestApi", Version: 1, Github: "https://github.com/krishpranav/gorestapi"},
}

var albums = []album{
	{ID: "1", Title: "ExampleOne", Artist: "User1", Price: 39.99},
	{ID: "2", Title: "ExampleTwo", Artist: "User2", Price: 39.99},
	{ID: "3", Title: "ExampleThree", Artist: "User3", Price: 39.99},
	{ID: "4", Title: "ExampleFour", Artist: "User4", Price: 39.99},
}

var mainpage = "Hello World"

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views/build", true)))

	api := router.Group("/api")
	{
		api.GET("/albums", getAlbums)
		router.GET("/albums/:id", getAlbumByID)
		router.POST("/albums", postAlbums)
		router.GET("/about", getAbout)
		router.GET("/", getMain)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed:", err)
	}

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAbout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, aboutapp)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
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

func getMain(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, mainpage)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
