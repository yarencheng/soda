package main

import (
	"encoding/json"
	"io/ioutil"
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
	File        string
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

	r.GET("/photo/:id/file", func(c *gin.Context) {
		ids := c.Param("id")
		log.Printf("ids = %v", ids)

		id, err := uuid.FromString(ids)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}
		log.Printf("id = %v", id)

		path := photoDir + "/photo/" + ids + "/data"
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		var photo Photo
		err = json.Unmarshal(b, &photo)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}
		log.Printf("photo = %#v", photo)

		dat, err := ioutil.ReadFile(photoDir + "/" + photo.File)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		c.Data(200, "", dat)
	})

	r.GET("/photo/:id", func(c *gin.Context) {

		ids := c.Param("id")
		log.Printf("ids = %v", ids)

		id, err := uuid.FromString(ids)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}
		log.Printf("id = %v", id)

		path := photoDir + "/" + ids + "/data"
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		var photo Photo
		err = json.Unmarshal(b, &photo)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "ok",
			"data":   photo,
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

		dir := photoDir + "/photo/" + photo.ID.String() + "/"
		log.Printf("dir = %v", dir)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		dst := dir + "/" + file.Filename
		log.Printf("dst = %v", dst)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			log.Printf("err: %v", err)
			c.AbortWithStatus(500)
			return
		}

		photo.File = "photo/" + photo.ID.String() + "/" + file.Filename
		jsonData, err := json.Marshal(photo)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		err = ioutil.WriteFile(dir+"/data", jsonData, 0644)
		if err != nil {
			log.Printf("err: %v", err)
			c.AbortWithStatus(500)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   photo,
		})
	})

	r.Run()
}
