package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

  func HealthCheck (c *fiber.Ctx) error{
  return c.SendString("Ok")
}

func main() {
	app := fiber.New()

	app.Get("/", HealthCheck)
	fmt.Println("hello world");
  log.Fatal(app.Listen(":3000"))
}
