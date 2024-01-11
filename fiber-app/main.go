package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users []User

func main() {
	app := fiber.New()

	app.Get("/api/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})
	app.Post("/api/users", func(c *fiber.Ctx) error {
		var NewUser User
		err := c.BodyParser(&NewUser)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Невозможно разобрать тело запроса"})
		}
		NewUser.ID = len(users) + 1
		users = append(users, NewUser)
		return c.Status(201).JSON(NewUser)
	})
	app.Get("/api/users/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Неверный формат ID"})
		}
		var founderUser User
		for _, v := range users {
			if v.ID == id {
				founderUser = v
				break
			}
		}
		if founderUser.ID == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "пользователь с таким ID не найден"})
		}
		return c.JSON(founderUser)
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
