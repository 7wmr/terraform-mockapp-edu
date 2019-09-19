package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"os/exec"
)

var msgEndpoint *string
var msgCredentials *string
var appPort *int

// failOnError - log an error messsage upon failure
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

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

// PostToQueue - post the request to RabbitMQ queue
func (r *Request) PostToQueue() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s@%s/", *msgCredentials, *msgEndpoint))

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Request", // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body, err := json.Marshal(r)
	failOnError(err, "Failed to parse to JSON")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
}

// SetHostname - run hostname command and set to property
func (r *Request) SetHostname() {
	cmd := exec.Command("hostname")
	stdout, err := cmd.Output()
	if err != nil {
		r.Hostname = "Unknown"
		return
	}
	r.Hostname = string(stdout)
}

// NewRequest - this will accept a new request.
func NewRequest(c *gin.Context) {
	var request Request
	request.SetUUID()
	request.SetHostname()
	c.JSON(http.StatusOK, request)
	request.PostToQueue()
}

func main() {
	appPort = flag.Int("port", 8080, "Web server port")
	msgEndpoint = flag.String("msg-endpoint", "", "RabbitMQ messaging endpoint")
	msgCredentials = flag.String("msg-credentials", "", "RabbitMQ messaging credentials")
	flag.Parse()

	router := gin.Default()
	router.LoadHTMLFiles("index.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Mock Application",
		})
	})

	v1 := router.Group("api/v1")
	{
		v1.GET("/info", NewRequest)
	}
	router.Run(fmt.Sprintf(":%d", *appPort))
}
