package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)


func NewRoom(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(http.StatusInternalServerError)
		return fmt.Errorf("no uuid has been received")
	}
	uuid, suuid, _ := createOrGetRoom(uuid)

	return nil
}

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

func RoomWebSocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}

	_, _, room := createOrGetRoom(uuid)

}

type Room struct {}

func createOrGetRoom(uuid string) (string, string, Room) {
	return "", "", Room{}
}