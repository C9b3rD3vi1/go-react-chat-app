
package main

import "fmt"

import (
    "log"
    //"net/http"
    
)

import (
    "github.com/gofiber/fiber/v2"
	//"github.com/gofiber/websocket/v2"
)



// WebSocket server
// Main chat application function
func main() {
	fmt.Println("Starting.....")
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!, Welcome to fiber chat server")
    })

    log.Fatal(app.Listen(":3000"))
}