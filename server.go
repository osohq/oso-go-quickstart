package main

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/osohq/go-oso"
)

func main() {
	app := fiber.New()
	oso, err := oso.NewOso()
	if err != nil {
		fmt.Printf("Failed to set up Oso: %v", err)
		return
	}

	oso.RegisterClass(reflect.TypeOf(Repository{}), nil)
	oso.RegisterClass(reflect.TypeOf(User{}), nil)
	if err := oso.LoadFiles([]string{"main.polar"}); err != nil {
		fmt.Printf("Failed to start: %s", err)
		return
	}

	app.Get("/repo/:repoName", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		repoName := c.Params("repoName")
		repository := GetRepositoryByName(repoName)
		err := oso.Authorize(GetCurrentUser(), "read", repository)
		if err == nil {
			return c.Status(200).SendString(fmt.Sprintf("<h1>A Repo</h1><p>Welcome to repo %s</p>", repository.Name))
		} else {
			return c.Status(404).SendString(fmt.Sprintf("<h1>Whoops!</h1><p>Repo named %s was not found</p>", repoName))
		}
	})
	if err := app.Listen(":5000"); err != nil {
		fmt.Printf("Failed to start: %s", err)
	}
}
