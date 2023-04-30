package controllers

import (
	"FiberProject/src/database"
	"FiberProject/src/middleware"
	"FiberProject/src/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}
	user := models.Player{
		FirsName: data["first_name"],
		LastName: data["last_name"],
		Email:    data["email"],
	}
	user.Password = user.SetPassword(data["password"])
	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "registering successfully",
		"data":    user,
	})
}
func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.Player

	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}
	if user.ComparePasswords(data["password"]) == false {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}
	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "login successfully",
	})
}

func User(c *fiber.Ctx) error {

	id, _ := middleware.GetUserId(c)
	var user models.Player
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func GetUserId(c *fiber.Ctx) {
	panic("unimplemented")
}

func Logout(c *fiber.Ctx) error {
	cookies := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookies)
	return c.JSON(fiber.Map{
		"message": "logout successfully",
	})
}
