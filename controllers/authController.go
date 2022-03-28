package controllers

import (
	"example/GO/database"
	"example/GO/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var responseStruct models.StatusResponse

func Register(c *fiber.Ctx) error {
	var userData map[string]string
	var databaseUser models.User

	if err := c.BodyParser(&userData); err != nil {
		return err
	}
	if userData["email"] == "" || userData["password"] == "" {
		responseStruct.Message = "internal server error"
		responseStruct.Status = fiber.StatusInternalServerError

		return c.JSON(responseStruct)
	}
	database.DB.Where("email =?", userData["email"]).First(&databaseUser)
	if databaseUser.Id != 0 {
		responseStruct.Message = "user have already existed"
		responseStruct.Status = fiber.StatusOK
		return c.JSON(responseStruct)
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(userData["password"]), 15)
	user := models.User{
		UserName: userData["username"],
		Email:    userData["email"],
		Password: password,
	}
	database.DB.Create(&user)
	return c.JSON(user)
}
func Login(c *fiber.Ctx) error {
	var userData map[string]string

	if err := c.BodyParser(&userData); err != nil {
		return err
	}
	if userData["email"] == "" || userData["password"] == "" {
		responseStruct.Message = "internal server error"
		responseStruct.Status = fiber.StatusInternalServerError

		return c.JSON(responseStruct)
	}
	var user models.User
	database.DB.Where("email =?", userData["email"]).First(&user)
	if user.Id == 0 {
		responseStruct.Message = "user not found"
		responseStruct.Status = fiber.StatusOK
		return c.JSON(responseStruct)
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(userData["password"])); err != nil {

		responseStruct.Message = "password incorrect"
		responseStruct.Status = fiber.StatusOK
		return c.JSON(responseStruct)

	}
	responseStruct.Message = "Login!"
	responseStruct.Status = fiber.StatusOK
	responseStruct.Data = map[string]interface{}{"token": user.Password}
	return c.JSON(responseStruct)

}
