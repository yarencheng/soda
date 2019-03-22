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

	r.DELETE("/photo/:id", func(c *gin.Context) {

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

		err = os.RemoveAll(photoDir + "/photo/" + id.String() + "/")
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/photo", func(c *gin.Context) {

		files, err := ioutil.ReadDir(photoDir + "/photo/")
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		photos := make([]Photo, len(files))

		for i, f := range files {
			log.Printf("f.Name() = %v", f.Name())

			b, err := ioutil.ReadFile(photoDir + "/photo/" + f.Name() + "/data")
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

			photos[i] = photo

		}

		c.JSON(200, gin.H{
			"status": "ok",
			"data":   photos,
		})
	})

	r.GET("/photo/:id/photo", func(c *gin.Context) {
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

	r.PUT("/photo/:id", func(c *gin.Context) {

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

		var oriPhoto Photo
		err = json.Unmarshal(b, &oriPhoto)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		var modPhoto Photo
		err = c.ShouldBindJSON(&modPhoto)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}
		log.Printf("photo = %#v", modPhoto)

		oriPhoto.Title = modPhoto.Title
		oriPhoto.Description = modPhoto.Description

		jsonData, err := json.Marshal(oriPhoto)
		if err != nil {
			log.Printf("err: %v", err)
			c.JSON(500, gin.H{
				"status": "failed",
			})
			return
		}

		err = ioutil.WriteFile(photoDir+"/photo/"+oriPhoto.ID.String()+"/data", jsonData, 0644)
		if err != nil {
			log.Printf("err: %v", err)
			c.AbortWithStatus(500)
			return
		}

		c.JSON(200, gin.H{
			"status": "ok",
			"data":   oriPhoto,
		})
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

		dst := dir + "/photo"
		log.Printf("dst = %v", dst)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			log.Printf("err: %v", err)
			c.AbortWithStatus(500)
			return
		}

		photo.File = "photo/" + photo.ID.String() + "/photo"
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
