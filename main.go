package main

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"os/exec"
	"strings"
)

// Request - declare structure
type Request struct {
	UUID     uuid.UUID `json:"uuid"`
	Hostname string    `json:"hostname"`
}

// SetUUID - generate a uuid for the Request
func (r *Request) SetUUID() {
	ID, _ := uuid.NewV4()
	r.UUID = ID
}

// SetHostname - run hostname command and set to property
func (r *Request) SetHostname() {
	cmd := exec.Command("hostname")
	stdout, err := cmd.Output()
	if err != nil {
		r.Hostname = "Unknown"
		return
	}
	r.Hostname = strings.ReplaceAll(string(stdout), "\n", "")
}

// NewRequest - this will accept a new request.
func NewRequest(c *gin.Context) {
	var request Request
	request.SetUUID()
	request.SetHostname()
	c.JSON(200, request)
}

func main() {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/info", NewRequest)
	}

	router.Run(":8080")
}
