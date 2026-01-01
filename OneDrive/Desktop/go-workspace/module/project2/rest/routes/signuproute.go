package routes

import (
	"errors"
	"example/my-project-go/module/project2/rest/models"
	"example/my-project-go/module/project2/rest/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignupRoute(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pass := user.Password
	user.Password = utils.HashPassword(pass)
	userID := user.Save()
	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully", "user_id": userID})
}

func LoginEvent(c *gin.Context) {
	var user models.User
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credentials invalid"})
		return
	}
	result, err := validateUsers(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credentials invalid"})
		return
	}
	fmt.Println("result:", result)
	jwtToken, err := utils.GenerateJWT(user.Email, int64(user.ID))
	if err != nil {
		fmt.Println("token is not generated please try again!!", err)
		c.JSON(http.StatusOK, gin.H{"Ok": "Login successfully!!!", "message": " token is not generated , please try after sometime!!!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Ok": "Login successfully!!!", "acesstoken": jwtToken})
}

func validateUsers(user models.User) (string, error) {
	users, err := user.GetUsersData()
	incomingPassWord := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	for _, userData := range users {
		if userData.Password == incomingPassWord && user.Email == userData.Email {
			return "ok", nil
		}
	}
	return "", errors.New("invalid password")
}
