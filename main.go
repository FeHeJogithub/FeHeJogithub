package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id`
	Title  string  `json:"title`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{

	{ID: "1", Title: "Blue Train", Artist: "John Contrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 50.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)

}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {

		return
	}
	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumsByID(c *gin.Context) {

	id := c.Param("id")

	for _, a := range albums {

		if a.ID == id {

			c.IndentedJSON(http.StatusOK, a)
			return

		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{

		"message": "album no encontrado"})

}
func putAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	for i, a := range albums {
		if a.ID == id {
			// Actualiza el álbum
			updatedAlbum.ID = id // Mantiene el ID original
			albums[i] = updatedAlbum

			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

func deleteAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			// Eliminar el álbum del slice
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "álbum eliminado"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

/*func deleteAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	for i, a := range albums {
		if a.ID == id {
			// Actualiza el álbum
			updatedAlbum.ID = id // Mantiene el ID original
			albums[i] = updatedAlbum

			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

/*
func putAlbumsByID(c *gin.Context) {

	id := c.Param("id")

	for _, a := range albums {

		if a.ID == id {

			c.IndentedJSON(http.StatusOK, a)
			return

		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{

		"message": "album no encontrado"})

}*/

func main() {

	//fmt.Println("Bienvenido a vinyl-api con fehejo")

	router := gin.Default()

	//router.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{

	//	"message": "Hola mundo con fehejo",
	//})
	//})
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", putAlbumsByID)
	router.DELETE("/albums/:id", deleteAlbumsByID)

	router.Run("localhost:8080")
}
