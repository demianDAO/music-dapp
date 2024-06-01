package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web3-music-platform/internal/api/models"
)

var musicCollection = []models.Music{}

func GetMusic(c *gin.Context) {
	c.JSON(http.StatusOK, musicCollection)
}

func GetMusicByID(c *gin.Context) {
	id := c.Param("id")
	for _, music := range musicCollection {
		if music.ID == id {
			c.JSON(http.StatusOK, music)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "music not found"})
}

func CreateMusic(c *gin.Context) {
	var newMusic models.Music
	if err := c.BindJSON(&newMusic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	musicCollection = append(musicCollection, newMusic)
	c.JSON(http.StatusCreated, newMusic)
}

func UpdateMusic(c *gin.Context) {
	id := c.Param("id")
	var updatedMusic models.Music
	if err := c.BindJSON(&updatedMusic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, music := range musicCollection {
		if music.ID == id {
			musicCollection[i] = updatedMusic
			c.JSON(http.StatusOK, updatedMusic)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "music not found"})
}

func DeleteMusic(c *gin.Context) {
	id := c.Param("id")
	for i, music := range musicCollection {
		if music.ID == id {
			musicCollection = append(musicCollection[:i], musicCollection[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "music deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "music not found"})
}
