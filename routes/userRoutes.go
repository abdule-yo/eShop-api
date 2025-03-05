package routes

import (
	"errors"

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
		userData := CreateResponseUser(user)
		responsesUsers = append(responsesUsers, userData)

	}

	return c.Status(200).JSON(fiber.Map{
		"Status":  200,
		"message": "Here are list of users",
		"users":   responsesUsers,
	})

}

// Todo: find a specific user for (get a user, update a user and delete a user)
func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("User is not found")
	}

	return nil
}

// Todo: return the user being found
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "There is no id passed",
		})
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Here is the user data",
		"user":    responseUser,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "There is no id passed",
		})
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Pleaes pass a data to update",
		})
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	return c.Status(200).JSON(fiber.Map{
		"status":  "200",
		"message": "User has been updated",
		"Data":    updateData,
	})

}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "There is no id passed",
		})
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := database.Database.Db.Delete(user, "id = ?", id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "200",
		"message": "User has been deleted",
	})

}
