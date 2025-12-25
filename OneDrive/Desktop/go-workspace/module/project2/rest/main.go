package main

import (
	"example/my-project-go/module/project2/rest/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world!!!")
	server := gin.Default()
	server.GET("/events", handleEvents)
	server.POST("/add-event", addEvent)
	server.DELETE("/delete-event", deleteEvent)
	server.PUT("/update-event", updateEvent)
	server.Run(":8080")
}

func handleEvents(c *gin.Context) {
	events := models.GetAllEvents()
	if events == nil {
		event := models.NewEvent(1, "Sample Event", "This is a sample event description",
			time.Now(), "Sample Location")
		models.AddEvent(*event)
		c.JSON(200, gin.H{"events": models.GetAllEvents()})
	} else {
		c.JSON(200, gin.H{"events": events})
	}
	fmt.Println("handle event is completed successfully")
}

func addEvent(c *gin.Context) {
	var event models.Event = models.Event{}
	//read event data from request body
	c.BindJSON(&event)
	models.AddEvent(event)
	c.JSON(200, gin.H{"Ok": "event is added successfully"})
	fmt.Println("event is added successfully")
}

func deleteEvent(c *gin.Context) {
	var event models.Event = models.Event{}
	//delete event data from request body
	c.BindJSON(&event)
	models.DeleteEvent(event)
	c.JSON(200, gin.H{"Ok": "event is deleted successfully"})
	fmt.Println("event is deleted successfully")
}

func updateEvent(c *gin.Context) {
	var event models.Event = models.Event{}
	//delete event data from request body
	c.BindJSON(&event)
	models.UpdateEvent(event)
	c.JSON(200, gin.H{"Ok": "event is updated successfully"})
}
