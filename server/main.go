package main

import (
	"context"
	// "fmt"
	"github.com/gin-contrib/cors"
	"log"
	"myproject/function"
	"myproject/model"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	// Set up a context and provide authentication options
	ctx := context.Background()
	// Replace "path/to/your/service-account-key.json" with the path to your service account key JSON file.
	// If running on Google Cloud, you can omit the option.WithCredentialsFile and the library will use Application Default Credentials.
	client, err := firestore.NewClient(ctx, "doan2-f82f0", option.WithCredentialsFile("credentials/doan2-f82f0-firebase-adminsdk-dufh3-d42255f4b5.json"))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		//   return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))
	r.POST("/login", func(c *gin.Context) {
		if err := function.Login(c, client); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": true,
			})
		}
	})
	r.GET("/updatestatus/:week/:studentid", func(c *gin.Context) {
		week := c.Param("week")
		studentid := c.Param("studentid")
		weekint, _ := strconv.Atoi(week)
		if err := function.UpdateStatus(c, client, weekint, studentid); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Already updated",
			})
		}
	})
	r.GET("/getAllStudent", func(c *gin.Context) {
		studentlist, err := function.GetAllStudent(c, client)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"data": studentlist,
			})
		}
	})
	r.GET("/getOneStudent/:id", func(c *gin.Context) {
		studentID := c.Param("id")
		student := function.GetOneStudent(c, client, studentID)
		if student == nil {
			c.JSON(400, gin.H{
				"error": "error",
			})
		} else {
			c.JSON(200, gin.H{
				"student": student,
			})
		}

	})
	r.POST("/addStudent", func(c *gin.Context) {
		var student model.Student
		c.ShouldBind(&student)
		if err := function.InsertStudent(c, client, &student); err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "insert succesfully!",
			})
		}

	})
	r.Run(":8080")
}
