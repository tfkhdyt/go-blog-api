package user

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	FindAllUsers(c *fiber.Ctx) error
	FindMyUser(c *fiber.Ctx) error
	FindOneUser(c *fiber.Ctx) error
	UpdateMyUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteMyUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}
