package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// TODO: Get all users
func GetUsers(c *fiber.Ctx) error {
	users := models.Users{}
	db.DataBase().Find(&users)
	return c.JSON(users)
}

// TODO: Get specific user
func GetUser(c *fiber.Ctx) error {
	if user, err := getUserById(c); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Resource not found")
	} else {
		return c.JSON(user)
	}

}

// TODO: Funcion to get a user in the database
func getUserById(c *fiber.Ctx) (models.User, *gorm.DB) {
	//obtener ID
	vars := c.Params("id")
	userId, _ := strconv.Atoi(vars)

	user := models.User{}
	if err := db.DataBase().First(&user, userId); err.Error != nil {
		return user, err
	} else {
		return user, nil
	}

}

// TODO: Create a user
func CreateUser(c *fiber.Ctx) error {
	//obtener registro
	user := models.User{}
	body := c.Body()

	if err := json.Unmarshal(body, &user); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Resource not found")
	} else {
		db.DataBase().Save(&user)
		c.Status(fiber.StatusCreated)
		return c.JSON(user)
	}
}

// TODO: Update specific user
func UpdateUser(c *fiber.Ctx) error {
	var userId int64

	if user_ant, err := getUserById(c); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Resource not found")
	} else {
		userId = user_ant.Id

		user := models.User{}
		body := c.Body()

		if err := json.Unmarshal(body, &user); err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "Resource not found")
		} else {
			user.Id = userId
			db.DataBase().Save(&user)
			return c.JSON(user)
		}
	}
}

// TODO: Delete specific user
func DeleteUser(c *fiber.Ctx) error {
	if user, err := getUserById(c); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Resource not found")
	} else {
		db.DataBase().Delete(&user)
		return c.JSON(user)
	}
}
