package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var photoDir string

type Photo struct {
	ID          uuid.UUID
	Title       string
	Description string
}

func main() {
	photoDir, exist := os.LookupEnv("PHOTO_DIR")
	if !exist {
		log.Fatalln("ENV PHOTO_DIR is not set")
	}
	log.Printf("photoDir = %v", photoDir)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/photo", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}
		log.Printf("file.Filename = %v", file.Filename)

		photo := &Photo{
			ID:          uuid.NewV4(),
			Title:       "empty",
			Description: "empty",
		}
		b, err := json.Marshal(photo)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		jsonData := string(b)
		log.Printf("jsonData = %v", jsonData)

		dir := photoDir + "/" + photo.ID.String() + "/"
		log.Printf("dir = %v", dir)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		dst := dir + "/photo"
		log.Printf("dst = %v", dst)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			log.Printf("err: %v", err)
			c.AbortWithStatus(500)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run()
}
