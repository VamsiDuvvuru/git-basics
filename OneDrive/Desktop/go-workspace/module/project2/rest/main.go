package main

import (
	"example/my-project-go/module/project2/rest/middleware"
	"example/my-project-go/module/project2/rest/models"
	"example/my-project-go/module/project2/rest/routes"
	"fmt"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDB()
	server := gin.Default()
	//using server group for authentication and authorization for the required apis
	//you can also use like server.post("/path/to/api" , middleware.Authenticate , eventHandler)
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.GET("/events", handleEvents)
	authenticated.POST("/add-event", addEvent)
	authenticated.DELETE("/delete/event/:id", deleteEvent)
	authenticated.PUT("/update/event/:id", updateEvent)
	authenticated.GET("/get/event/:id", getEventById)

	server.POST("/user/signup", routes.SignupRoute)
	server.POST("/login", routes.LoginEvent)
	server.Run(":8080")
}

func handleEvents(c *gin.Context) {
	userId, exists := c.Get("userId")
	events := models.GetAllEvents()
	if events == nil {
		event := models.NewEvent(1, "Sample Event", "This is a sample event description",
			time.Now(), "Sample Location")
		models.AddEvent(*event)
		c.JSON(200, gin.H{"events": models.GetAllEvents()})
		return
	} else {
		if exists {
			c.JSON(200, gin.H{"events": events, "userId": userId})
		} else {
			c.JSON(200, gin.H{"events": events})
		}
		return
	}
}

func addEvent(c *gin.Context) {
	userId, exists := c.Get("userId")
	var event models.Event = models.Event{}
	//read event data from request body
	c.BindJSON(&event)
	models.AddEvent(event)
	if exists {
		c.JSON(200, gin.H{"Ok": "event is added successfully", "userId ": userId})
		return
	}
	c.JSON(200, gin.H{"Ok": "event is added successfully"})
	fmt.Println("event is added successfully")
}

func deleteEvent(c *gin.Context) {
	userId, exists := c.Get("userId")
	id := c.Param("id")
	//convert string to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "id parameter is required"})
		return
	}

	if models.GetEventByID(intID) == nil {
		c.JSON(404, gin.H{"error": "event not found, id is" + strconv.Itoa(intID)})
		return
	}
	models.DeleteEvent(intID)
	if exists {
		c.JSON(200, gin.H{"Ok": "event is deleted successfully", "userId ": userId})
		return
	}
	c.JSON(200, gin.H{"Ok": "event is deleted successfully"})
}

func updateEvent(c *gin.Context) {
	userId, exists := c.Get("userId")
	id := c.Param("id")
	//convert string to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "id parameter is required"})
		return
	}
	var event models.Event = models.Event{}
	//delete event data from request body
	c.BindJSON(&event)
	if event.ID != intID {
		c.JSON(500, gin.H{"error": "event ID in body does not match URL parameter"})
		return
	}
	if models.GetEventByID(event.ID) == nil {
		c.JSON(404, gin.H{"error": "event not found"})
		return
	}
	models.UpdateEvent(event)
	if exists {
		c.JSON(200, gin.H{"Ok": "event is updated successfully", "userId ": userId})
		return
	}
	c.JSON(200, gin.H{"Ok": "event is updated successfully"})
}

func getEventById(c *gin.Context) {
	userId, exists := c.Get("userId")
	id := c.Param("id")
	//cpnvert string to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "id parameter is required"})
		return
	}
	event := models.GetEventByID(intID)
	if event == nil {
		c.JSON(404, gin.H{"error": "event not found"})
		return
	}
	if exists {
		c.JSON(200, gin.H{"event": event, "userId ": userId})
		return
	}
	c.JSON(200, gin.H{"event": event})
}
