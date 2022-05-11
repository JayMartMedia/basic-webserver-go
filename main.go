package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID		 string	 `json:"id"`
	Title	 string	 `json:"title"`
	Artist string  `json:"artist"`
	Price	 float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbumById)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumById gets a single album with the id provided in the request
// and retuns that album
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// loop over albums, look for matching id to return
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	// if no matches are found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album
	
	// TODO: add logic to check whether all required fields are present
	// TODO: add logic to check whether an album with that id has already been created
	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice (writes to top level albums slice)
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// deleteAlbumById removes the albums with the provided id
func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	// loop over albums, look for matching id to remove
	for i, album := range albums {
		if album.ID == id {
			// Removes the album from the slice (writes to top level albums slice)
			albums = removeAlbum(albums, i)
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	// if no matches are found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Removes a single album at the provided index from a slice of albums
func removeAlbum(slice []album, index int) []album {
	return append(slice[:index], slice[index+1:]...)
}