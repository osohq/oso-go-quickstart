package main

import (
	"fmt"
	"reflect"
	"strconv"

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

	if err := oso.RegisterClass(reflect.TypeOf(Repository{}), nil); err != nil {
		fmt.Printf("Failed to start: %s", err)
		return
	}
	if err := oso.RegisterClass(reflect.TypeOf(User{}), nil); err != nil {
		fmt.Printf("Failed to start: %s", err)
		return
	}
	if err := oso.LoadFile("main.polar"); err != nil {
		fmt.Printf("Failed to start: %s", err)
		return
	}

	app.Get("/repo/:repoId", func(c *fiber.Ctx) error {
		repoId, err := strconv.Atoi(c.Params("repoId"))
		if err != nil {
			return c.SendStatus(400)
		}
		repository := 
		allowed, err := oso.IsAllowed(GetCurrentUser(), "read", GetRepositoryById(repoId))
		if err == nil && allowed {
			return c.Status(200).SendString(fmt.Sprintf("<h1>A Repo</h1><p>Welcome to repo %s</p>", repository.Name))
		} else {
			return c.Status(404).SendString("<h1>Whoops!</h1><p>That repo was not found</p>")

		}
	})
	if err := app.Listen(":5000"); err != nil {
		fmt.Printf("Failed to start: %s", err)
	}
}
