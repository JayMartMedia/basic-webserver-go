package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jaymartmedia/basic_webserver_go/album"
)

var albums = make([]album.Album, 0)

func seedData() {
	albums = append(albums, album.New("Blue Train", "John Coltrane", 56.99))
	albums = append(albums, album.New("Jeru", "Gerry Mulligan", 17.99))
	albums = append(albums, album.New("Sarah Vaughan and Clifford Brown", "Sarah Vaughan", 39.99))
}

func main() {
	seedData()

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.PUT("/albums/:id", updateAlbumById)

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

// postAlbum adds an album from JSON received in the request body
func postAlbum(c *gin.Context) {
	var newAlbum album.Album
	
	// TODO: add logic to check whether all required fields are present
	// TODO: add logic to check whether an album with that id has already been created
	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	var a = album.New(newAlbum.Title,newAlbum.Artist,newAlbum.Price);

	// Add the new album to the slice (writes to top level albums slice)\
	albums = append(albums, a)
	c.IndentedJSON(http.StatusCreated, a)
}

// deleteAlbumById removes the album with the provided id
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

// updateAlbum updates the album with the provided id
func updateAlbumById(c *gin.Context) {
	id := c.Param("id")

	var newAlbum album.Album
	
	// TODO: add logic to check whether all required fields are present
	// TODO: add logic to check whether an album with that id has already been created
	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// loop over albums, look for matching id to overwrite
	for i, album := range albums {
		if album.ID == id {
			albums[i] = newAlbum
			c.IndentedJSON(http.StatusOK, newAlbum)
			return
		}
	}

	// if no matches are found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Removes a single album at the provided index from a slice of albums
func removeAlbum(slice []album.Album, index int) []album.Album {
	return append(slice[:index], slice[index+1:]...)
}