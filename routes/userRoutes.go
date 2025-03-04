package routes

import (
	"github.com/abdule-yo/eCommerce-api/database"
	"github.com/abdule-yo/eCommerce-api/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body, try again",
		})
	}

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	responsesUsers := []User{}

	for _, user := range users {
		responsesUser := CreateResponseUser(user)
		responsesUsers = append(responsesUsers, responsesUser)

	}

	return c.Status(200).JSON(fiber.Map{
		"Status":  200,
		"message": "Here are list of users",
		"users":   responsesUsers,
	})

}
